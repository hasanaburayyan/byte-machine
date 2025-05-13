import * as vscode from 'vscode';
import * as Net from 'net';
import { ByteMachineDebugSession } from './adapter';

export function activate(context: vscode.ExtensionContext) {
  console.log("ACTIVATING!!");
  const server = Net.createServer(socket => {
    console.log("Inside create server!")
    const session = new ByteMachineDebugSession();
    session.setRunAsServer(true);
    session.start(socket, socket);
  });

  server.listen(0, () => {
    const port = (server.address() as Net.AddressInfo).port
    context.subscriptions.push(
      vscode.debug.registerDebugConfigurationProvider("bytemachine", {
        resolveDebugConfiguration(folder, config) {
          config.debugServer = port;
          return config;
        }
      })
    )
  });
}