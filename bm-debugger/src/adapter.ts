import {
  LoggingDebugSession,
  InitializedEvent,
  StoppedEvent,
  TerminatedEvent,
  Breakpoint,
} from 'vscode-debugadapter';
import { DebugProtocol } from 'vscode-debugprotocol';
import { spawn, ChildProcess } from 'child_process';
import * as fs from 'fs';

interface LaunchRequestArgs extends DebugProtocol.LaunchRequestArguments {
  program: string;
  sourcemap: string;
}

export class ByteMachineDebugSession extends LoggingDebugSession {
  private static threadID = 1;
  private vmProcess: ChildProcess | null = null;
  private sourcemap: Record<string, number> = {};
  private ip = 0;

  public constructor() {
    super("byte-debug.txt");
    console.log("CONSTRUCTOR")
    this.setDebuggerLinesStartAt1(true);
    this.setDebuggerColumnsStartAt1(true);
  }

  protected initializeRequest(response: DebugProtocol.InitializeResponse): void {
    console.log("INITIALIZED!")
    response.body = {
      supportsConfigurationDoneRequest: true,
    };
    this.sendResponse(response);
    this.sendEvent(new InitializedEvent());
  }

  protected launchRequest(response: DebugProtocol.LaunchResponse, args: LaunchRequestArgs): void {
    this.sourcemap = JSON.parse(fs.readFileSync(args.sourcemap, 'utf8')).map;
    console.log('Spawning with:', args);
    this.vmProcess = spawn('bmdebug', ['--debug', '--input', args.program]);
    this.vmProcess.stdout?.on('data', (data) => {
      console.log("RECEIVED DATA: ", data.toString())
      try {
        const event = JSON.parse(data.toString());
        if (event.event === 'stopped') {
          this.ip = event.ip;
          this.sendEvent(new StoppedEvent('breakpoint', ByteMachineDebugSession.threadID));
        } else if (event.event === 'terminated') {
          this.sendEvent(new TerminatedEvent());
        }
      } catch (e) {
        console.error('Invalid JSON from VM:', data.toString());
      }
    });
    console.log("Spawn done")

    this.sendResponse(response);
  }

  protected setBreakPointsRequest(response: DebugProtocol.SetBreakpointsResponse, args: DebugProtocol.SetBreakpointsArguments): void {
    const breakpoints = (args.breakpoints || []).map(bp => new Breakpoint(true, bp.line));
    console.log("BREAKPONTS:", args.breakpoints)
    response.body = { breakpoints };
    this.sendResponse(response);
  }

  protected stackTraceRequest(response: DebugProtocol.StackTraceResponse): void {
    const line = this.sourcemap[this.ip.toString()] || 1;
    response.body = {
      stackFrames: [{
        id: 1,
        name: 'main',
        line: line,
        column: 1,
        source: { name: 'sample.bm', path: 'sample.bm' },
      }],
      totalFrames: 1,
    };
    this.sendResponse(response);
  }

  protected continueRequest(response: DebugProtocol.ContinueResponse): void {
    this.vmProcess?.stdin?.write('{"command": "continue"}\n');
    this.sendResponse(response);
  }

  protected nextRequest(response: DebugProtocol.NextResponse): void {
    this.vmProcess?.stdin?.write('{"command": "step"}\n');
    this.sendResponse(response);
  }
}
