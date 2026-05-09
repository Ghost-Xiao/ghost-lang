import {
	createConnection,
	TextDocuments,
	Diagnostic,
	DiagnosticSeverity,
	ProposedFeatures,
	InitializeParams,
	DidChangeConfigurationNotification,
	CompletionItem,
	CompletionItemKind,
	TextDocumentPositionParams,
	TextDocumentSyncKind,
	InitializeResult
} from 'vscode-languageserver/node';

import {
	TextDocument
} from 'vscode-languageserver-textdocument';

// 为服务器创建连接，使用 Node 的 IPC 作为传输
// 同时包含所有预览/提议的 LSP 功能
const connection = createConnection(ProposedFeatures.all);

// 创建一个简单的文本文档管理器
const documents: TextDocuments<TextDocument> = new TextDocuments(TextDocument);

let hasConfigurationCapability = false;
let hasWorkspaceFolderCapability = false;
let hasDiagnosticRelatedInformationCapability = false;

connection.onInitialize((params: InitializeParams) => {
	const capabilities = params.capabilities;

	// 客户端是否支持 `workspace/configuration` 请求？
	// 如果不支持，我们会回退使用全局设置
	hasConfigurationCapability = !!(
		capabilities.workspace && !!capabilities.workspace.configuration
	);
	hasWorkspaceFolderCapability = !!(
		capabilities.workspace && !!capabilities.workspace.workspaceFolders
	);
	hasDiagnosticRelatedInformationCapability = !!(
		capabilities.textDocument &&
		capabilities.textDocument.publishDiagnostics &&
		capabilities.textDocument.publishDiagnostics.relatedInformation
	);

	const result: InitializeResult = {
		capabilities: {
			textDocumentSync: {
				openClose: true,
				change: TextDocumentSyncKind.Incremental,
				// 保存时发送完整文档内容
				save: {
					includeText: true
				}
			},
			completionProvider: {
				resolveProvider: true
			}
		}
	};
	if (hasWorkspaceFolderCapability) {
		result.capabilities.workspace = {
			workspaceFolders: {
				supported: true
			}
		};
	}
	return result;
});

connection.onInitialized(() => {
	if (hasConfigurationCapability) {
		// 注册所有配置变更
		connection.client.register(DidChangeConfigurationNotification.type, undefined);
	}
	if (hasWorkspaceFolderCapability) {
		connection.workspace.onDidChangeWorkspaceFolders(_event => {
			connection.console.log('收到工作区文件夹变更事件');
		});
	}
});

// 示例设置
interface GhostSettings {
	maxNumberOfProblems: number;
}

// 全局设置，当客户端不支持 `workspace/configuration` 请求时使用
// 请注意，当使用此示例提供的客户端时不会发生这种情况，但可能会在其他客户端发生
const defaultSettings: GhostSettings = { maxNumberOfProblems: 100 };
let globalSettings: GhostSettings = defaultSettings;

// 缓存所有打开文档的设置
const documentSettings: Map<string, Thenable<GhostSettings>> = new Map();

connection.onDidChangeConfiguration(change => {
	if (hasConfigurationCapability) {
		// 重置所有缓存的文档设置
		documentSettings.clear();
	} else {
		globalSettings = <GhostSettings>(
			(change.settings.ghost || defaultSettings)
		);
	}

	// 重新验证所有打开的文本文档
	documents.all().forEach(validateGhostDocument);
});

function getDocumentSettings(resource: string): Thenable<GhostSettings> {
	if (!hasConfigurationCapability) {
		return Promise.resolve(globalSettings);
	}
	let result = documentSettings.get(resource);
	if (!result) {
		result = connection.workspace.getConfiguration({
			scopeUri: resource,
			section: 'ghost'
		});
		documentSettings.set(resource, result);
	}
	return result;
}

// 仅保留打开文档的设置
documents.onDidClose(e => {
	documentSettings.delete(e.document.uri);
});

// 文本文档的内容已更改，此事件在首次打开或内容更改时触发
documents.onDidChangeContent(change => {
	validateGhostDocument(change.document);
});

async function validateGhostDocument(textDocument: TextDocument): Promise<void> {
	// 在这个简单示例中，我们为每次验证运行获取设置
	const settings = await getDocumentSettings(textDocument.uri);

	const text = textDocument.getText();
	const diagnostics: Diagnostic[] = [];
	let problems = 0;

	// 检查常见错误
	// 1. 未闭合的字符串字面量
	let i = 0;
	while (i < text.length && problems < settings.maxNumberOfProblems) {
		// 检查字符串
		if (text[i] === '"' || text[i] === "'" || text[i] === '`') {
			const quote = text[i];
			const start = i;
			i++;
			let found = false;
			while (i < text.length) {
				if (text[i] === '\\') {
					i += 2;
				} else if (text[i] === quote) {
					found = true;
					i++;
					break;
				} else {
					i++;
				}
			}
			if (!found) {
				const diagnostic: Diagnostic = {
					severity: DiagnosticSeverity.Error,
					range: {
						start: textDocument.positionAt(start),
						end: textDocument.positionAt(text.length)
					},
					message: `未闭合的${quote === '"' ? '双' : quote === "'" ? '单' : '反引号'}字符串字面量`,
					source: 'ghost'
				};
				diagnostics.push(diagnostic);
				problems++;
			}
		} else {
			i++;
		}
	}

	// 2. 不匹配的括号
	const brackets: { [key: string]: string } = { '(': ')', '{': '}', '[': ']' };
	const stack: { type: string, pos: number }[] = [];
	i = 0;
	while (i < text.length && problems < settings.maxNumberOfProblems) {
		const char = text[i];
		if (brackets[char]) {
			stack.push({ type: char, pos: i });
		} else if (Object.values(brackets).includes(char)) {
			const matching = Object.keys(brackets).find(key => brackets[key] === char);
			if (stack.length > 0 && stack[stack.length - 1].type === matching) {
				stack.pop();
			} else {
				const diagnostic: Diagnostic = {
					severity: DiagnosticSeverity.Error,
					range: {
						start: textDocument.positionAt(i),
						end: textDocument.positionAt(i + 1)
					},
					message: `不匹配的闭合括号 '${char}'`,
					source: 'ghost'
				};
				diagnostics.push(diagnostic);
				problems++;
			}
		}
		i++;
	}

	for (const item of stack) {
		if (problems >= settings.maxNumberOfProblems) break;
		const diagnostic: Diagnostic = {
			severity: DiagnosticSeverity.Error,
			range: {
				start: textDocument.positionAt(item.pos),
				end: textDocument.positionAt(item.pos + 1)
			},
			message: `未闭合的 '${item.type}'，需要 '${brackets[item.type]}'`,
			source: 'ghost'
		};
		diagnostics.push(diagnostic);
		problems++;
	}

	// 向 VS Code 发送计算出的诊断信息
	connection.sendDiagnostics({ uri: textDocument.uri, diagnostics });
}

connection.onDidChangeWatchedFiles(_change => {
	// VS Code 中被监控的文件发生了变更
	connection.console.log('收到文件变更事件');
});

// 此处理程序提供补全项的初始列表
connection.onCompletion(
	(_textDocumentPosition: TextDocumentPositionParams): CompletionItem[] => {
		// 传递的参数包含请求代码补全的文本文档位置
		// 在此示例中我们忽略此信息，始终提供相同的补全项

		// Ghost 语言关键字
		const keywords = [
			{ label: 'var', kind: CompletionItemKind.Keyword, documentation: '声明变量' },
			{ label: 'const', kind: CompletionItemKind.Keyword, documentation: '声明常量' },
			{ label: 'func', kind: CompletionItemKind.Keyword, documentation: '定义函数' },
			{ label: 'if', kind: CompletionItemKind.Keyword, documentation: '条件语句' },
			{ label: 'else', kind: CompletionItemKind.Keyword, documentation: '条件语句的 else 分支' },
			{ label: 'for', kind: CompletionItemKind.Keyword, documentation: 'for 循环' },
			{ label: 'foreach', kind: CompletionItemKind.Keyword, documentation: 'foreach 循环' },
			{ label: 'in', kind: CompletionItemKind.Keyword, documentation: '范围运算符' },
			{ label: 'step', kind: CompletionItemKind.Keyword, documentation: '循环步长' },
			{ label: 'return', kind: CompletionItemKind.Keyword, documentation: '从函数返回' },
			{ label: 'break', kind: CompletionItemKind.Keyword, documentation: '跳出循环' },
			{ label: 'continue', kind: CompletionItemKind.Keyword, documentation: '继续下一次循环' },
			{ label: 'namespace', kind: CompletionItemKind.Keyword, documentation: '定义命名空间' },
			{ label: 'import', kind: CompletionItemKind.Keyword, documentation: '导入模块' },
			{ label: 'class', kind: CompletionItemKind.Keyword, documentation: '定义类' },
			{ label: 'extends', kind: CompletionItemKind.Keyword, documentation: '继承自类' },
			{ label: 'this', kind: CompletionItemKind.Keyword, documentation: '引用当前实例' },
			{ label: 'super', kind: CompletionItemKind.Keyword, documentation: '引用父类' },
			{ label: 'try', kind: CompletionItemKind.Keyword, documentation: 'try 块' },
			{ label: 'catch', kind: CompletionItemKind.Keyword, documentation: '捕获异常' },
			{ label: 'finally', kind: CompletionItemKind.Keyword, documentation: 'finally 块' },
			{ label: 'throw', kind: CompletionItemKind.Keyword, documentation: '抛出异常' },
			{ label: 'true', kind: CompletionItemKind.Value, documentation: '布尔值 true' },
			{ label: 'false', kind: CompletionItemKind.Value, documentation: '布尔值 false' },
			{ label: 'null', kind: CompletionItemKind.Value, documentation: '空值' },
			{ label: 'contains', kind: CompletionItemKind.Operator, documentation: '包含运算符' }
		];

		// 内置函数
		const builtins = [
			{ label: 'print', kind: CompletionItemKind.Function, documentation: '打印值，不换行' },
			{ label: 'println', kind: CompletionItemKind.Function, documentation: '打印值，换行' },
			{ label: 'input', kind: CompletionItemKind.Function, documentation: '读取用户输入' },
			{ label: 'len', kind: CompletionItemKind.Function, documentation: '获取字符串、列表或映射的长度' },
			{ label: 'power', kind: CompletionItemKind.Function, documentation: '计算数字的幂' }
		];

		// math 模块函数
		const mathFuncs = [
			{ label: 'math.factorial', kind: CompletionItemKind.Function, documentation: '计算阶乘 n!' },
			{ label: 'math.comb', kind: CompletionItemKind.Function, documentation: '计算组合数 C(n, k)' },
			{ label: 'math.perm', kind: CompletionItemKind.Function, documentation: '计算排列数 P(n, k)' },
			{ label: 'math.abs', kind: CompletionItemKind.Function, documentation: '绝对值' },
			{ label: 'math.sqrt', kind: CompletionItemKind.Function, documentation: '平方根' },
			{ label: 'math.sin', kind: CompletionItemKind.Function, documentation: '正弦' },
			{ label: 'math.cos', kind: CompletionItemKind.Function, documentation: '余弦' },
			{ label: 'math.tan', kind: CompletionItemKind.Function, documentation: '正切' },
			{ label: 'math.asin', kind: CompletionItemKind.Function, documentation: '反正弦' },
			{ label: 'math.acos', kind: CompletionItemKind.Function, documentation: '反余弦' },
			{ label: 'math.atan', kind: CompletionItemKind.Function, documentation: '反正切' },
			{ label: 'math.log', kind: CompletionItemKind.Function, documentation: '对数' },
			{ label: 'math.lg', kind: CompletionItemKind.Function, documentation: '常用对数' },
			{ label: 'math.ln', kind: CompletionItemKind.Function, documentation: '自然对数' },
			{ label: 'math.floor', kind: CompletionItemKind.Function, documentation: '向下取整' },
			{ label: 'math.ceil', kind: CompletionItemKind.Function, documentation: '向上取整' },
			{ label: 'math.round', kind: CompletionItemKind.Function, documentation: '四舍五入' },
			{ label: 'math.min', kind: CompletionItemKind.Function, documentation: '最小值' },
			{ label: 'math.max', kind: CompletionItemKind.Function, documentation: '最大值' },
			{ label: 'math.sum', kind: CompletionItemKind.Function, documentation: '求和' },
			{ label: 'math.product', kind: CompletionItemKind.Function, documentation: '乘积' },
			{ label: 'math.mean', kind: CompletionItemKind.Function, documentation: '平均值' },
			{ label: 'math.median', kind: CompletionItemKind.Function, documentation: '中位数' },
			{ label: 'math.variance', kind: CompletionItemKind.Function, documentation: '方差' },
			{ label: 'math.stdDev', kind: CompletionItemKind.Function, documentation: '标准差' },
			{ label: 'math.rand', kind: CompletionItemKind.Function, documentation: '随机数' },
			{ label: 'math.randInt', kind: CompletionItemKind.Function, documentation: '随机整数' },
			{ label: 'math.PI', kind: CompletionItemKind.Constant, documentation: '圆周率 π' },
			{ label: 'math.E', kind: CompletionItemKind.Constant, documentation: '自然对数底 e' },
			{ label: 'math.TAU', kind: CompletionItemKind.Constant, documentation: '2π' }
		];

		// 内置错误类
		const errors = [
			{ label: 'Error', kind: CompletionItemKind.Class, documentation: '基础错误类' },
			{ label: 'OperationError', kind: CompletionItemKind.Class, documentation: '操作错误' },
			{ label: 'MathError', kind: CompletionItemKind.Class, documentation: '数学错误' },
			{ label: 'TypeError', kind: CompletionItemKind.Class, documentation: '类型错误' },
			{ label: 'IndexError', kind: CompletionItemKind.Class, documentation: '索引错误' },
			{ label: 'VariableError', kind: CompletionItemKind.Class, documentation: '变量错误' },
			{ label: 'ArgumentError', kind: CompletionItemKind.Class, documentation: '参数错误' },
			{ label: 'ModuleError', kind: CompletionItemKind.Class, documentation: '模块错误' }
		];

		// 内置模块
		const modules = [
			{ label: '"math"', kind: CompletionItemKind.Module, documentation: '数学模块' },
			{ label: '"fmt"', kind: CompletionItemKind.Module, documentation: '格式化模块' },
			{ label: '"io"', kind: CompletionItemKind.Module, documentation: 'I/O 模块' },
			{ label: '"time"', kind: CompletionItemKind.Module, documentation: '时间模块' }
		];

		return [
			...keywords.map(item => ({
				label: item.label,
				kind: item.kind,
				data: item,
				documentation: item.documentation
			})),
			...builtins.map(item => ({
				label: item.label,
				kind: item.kind,
				data: item,
				documentation: item.documentation
			})),
			...mathFuncs.map(item => ({
				label: item.label,
				kind: item.kind,
				data: item,
				documentation: item.documentation
			})),
			...errors.map(item => ({
				label: item.label,
				kind: item.kind,
				data: item,
				documentation: item.documentation
			})),
			...modules.map(item => ({
				label: item.label,
				kind: item.kind,
				data: item,
				documentation: item.documentation
			}))
		];
	}
);

// 此处理程序为补全列表中选择的项目解析附加信息
connection.onCompletionResolve(
	(item: CompletionItem): CompletionItem => {
		// 如果需要，我们可以在这里丰富补全项
		return item;
	}
);

// 让文本文档管理器在连接上监听
// 打开、更改和关闭文本文档事件
documents.listen(connection);

// 监听连接
connection.listen();
