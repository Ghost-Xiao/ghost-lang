import * as path from 'path';
import { workspace, ExtensionContext } from 'vscode';

import {
	LanguageClient,
	LanguageClientOptions,
	ServerOptions,
	TransportKind
} from 'vscode-languageclient/node';

let client: LanguageClient;

export function activate(context: ExtensionContext) {
	// 服务器在 node 中实现
	const serverModule = context.asAbsolutePath(
		path.join('out', 'server.js')
	);
	// 服务器的调试选项
	// --inspect=6009: 在 Node 的检查器模式运行服务器，以便 VS Code 可以附加到服务器进行调试
	const debugOptions = { execArgv: ['--nolazy', '--inspect=6009'] };

	// 如果扩展在调试模式下启动，则使用调试服务器选项
	// 否则使用运行选项
	const serverOptions: ServerOptions = {
		run: { module: serverModule, transport: TransportKind.ipc },
		debug: {
			module: serverModule,
			transport: TransportKind.ipc,
			options: debugOptions
		}
	};

	// 控制语言客户端的选项
	const clientOptions: LanguageClientOptions = {
		// 为纯文本文档注册服务器
		documentSelector: [{ scheme: 'file', language: 'ghost' }],
		synchronize: {
			// 向服务器通知工作区中 .clientrc 文件的变更
			fileEvents: workspace.createFileSystemWatcher('**/.clientrc')
		}
	};

	// 创建语言客户端并启动客户端
	client = new LanguageClient(
		'ghostLanguageServer',
		'Ghost 语言服务器',
		serverOptions,
		clientOptions
	);

	// 启动客户端，这也将同时启动服务器
	client.start();
}

export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}
	return client.stop();
}
