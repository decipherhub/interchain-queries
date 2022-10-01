/* eslint-disable */
import {
  QueryResult,
  queryResultFromJSON,
  queryResultToJSON,
} from "../../ibc-query/v1/crosschainquery";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Height } from "../../ibc/core/client/v1/client";
import { ProofSpec } from "../../proofs";

export const protobufPackage = "interchain_queries.ibc_query.v1";

/** MsgSubmitCrossChainQuery */
export interface MsgSubmitCrossChainQuery {
  id: string;
  path: string;
  localTimeoutHeight: Height | undefined;
  localTimeoutStamp: number;
  queryHeight: number;
  clientId: string;
  /** sender address */
  sender: string;
  sourcePort: string;
  sourceChannel: string;
}

/** MsgSubmitCrossChainQueryResponse */
export interface MsgSubmitCrossChainQueryResponse {
  queryId: string;
  capKey: number;
}

/** MsgSubmitCrossChainQueryResult */
export interface MsgSubmitCrossChainQueryResult {
  id: string;
  path: string;
  queryHeight: number;
  result: QueryResult;
  data: Uint8Array;
  sender: string;
  /** TODO: Proof specifications used in verifying counterparty state */
  proofSpecs: ProofSpec[];
}

/** MsgSubmitCrossChainQueryResultResponse */
export interface MsgSubmitCrossChainQueryResultResponse {}

const baseMsgSubmitCrossChainQuery: object = {
  id: "",
  path: "",
  localTimeoutStamp: 0,
  queryHeight: 0,
  clientId: "",
  sender: "",
  sourcePort: "",
  sourceChannel: "",
};

export const MsgSubmitCrossChainQuery = {
  encode(
    message: MsgSubmitCrossChainQuery,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.path !== "") {
      writer.uint32(18).string(message.path);
    }
    if (message.localTimeoutHeight !== undefined) {
      Height.encode(
        message.localTimeoutHeight,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.localTimeoutStamp !== 0) {
      writer.uint32(32).uint64(message.localTimeoutStamp);
    }
    if (message.queryHeight !== 0) {
      writer.uint32(40).uint64(message.queryHeight);
    }
    if (message.clientId !== "") {
      writer.uint32(50).string(message.clientId);
    }
    if (message.sender !== "") {
      writer.uint32(58).string(message.sender);
    }
    if (message.sourcePort !== "") {
      writer.uint32(66).string(message.sourcePort);
    }
    if (message.sourceChannel !== "") {
      writer.uint32(74).string(message.sourceChannel);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCrossChainQuery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCrossChainQuery,
    } as MsgSubmitCrossChainQuery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.path = reader.string();
          break;
        case 3:
          message.localTimeoutHeight = Height.decode(reader, reader.uint32());
          break;
        case 4:
          message.localTimeoutStamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.queryHeight = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.clientId = reader.string();
          break;
        case 7:
          message.sender = reader.string();
          break;
        case 8:
          message.sourcePort = reader.string();
          break;
        case 9:
          message.sourceChannel = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitCrossChainQuery {
    const message = {
      ...baseMsgSubmitCrossChainQuery,
    } as MsgSubmitCrossChainQuery;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (
      object.localTimeoutHeight !== undefined &&
      object.localTimeoutHeight !== null
    ) {
      message.localTimeoutHeight = Height.fromJSON(object.localTimeoutHeight);
    } else {
      message.localTimeoutHeight = undefined;
    }
    if (
      object.localTimeoutStamp !== undefined &&
      object.localTimeoutStamp !== null
    ) {
      message.localTimeoutStamp = Number(object.localTimeoutStamp);
    } else {
      message.localTimeoutStamp = 0;
    }
    if (object.queryHeight !== undefined && object.queryHeight !== null) {
      message.queryHeight = Number(object.queryHeight);
    } else {
      message.queryHeight = 0;
    }
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = String(object.clientId);
    } else {
      message.clientId = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.sourcePort !== undefined && object.sourcePort !== null) {
      message.sourcePort = String(object.sourcePort);
    } else {
      message.sourcePort = "";
    }
    if (object.sourceChannel !== undefined && object.sourceChannel !== null) {
      message.sourceChannel = String(object.sourceChannel);
    } else {
      message.sourceChannel = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQuery): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.path !== undefined && (obj.path = message.path);
    message.localTimeoutHeight !== undefined &&
      (obj.localTimeoutHeight = message.localTimeoutHeight
        ? Height.toJSON(message.localTimeoutHeight)
        : undefined);
    message.localTimeoutStamp !== undefined &&
      (obj.localTimeoutStamp = message.localTimeoutStamp);
    message.queryHeight !== undefined &&
      (obj.queryHeight = message.queryHeight);
    message.clientId !== undefined && (obj.clientId = message.clientId);
    message.sender !== undefined && (obj.sender = message.sender);
    message.sourcePort !== undefined && (obj.sourcePort = message.sourcePort);
    message.sourceChannel !== undefined &&
      (obj.sourceChannel = message.sourceChannel);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQuery>
  ): MsgSubmitCrossChainQuery {
    const message = {
      ...baseMsgSubmitCrossChainQuery,
    } as MsgSubmitCrossChainQuery;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (
      object.localTimeoutHeight !== undefined &&
      object.localTimeoutHeight !== null
    ) {
      message.localTimeoutHeight = Height.fromPartial(
        object.localTimeoutHeight
      );
    } else {
      message.localTimeoutHeight = undefined;
    }
    if (
      object.localTimeoutStamp !== undefined &&
      object.localTimeoutStamp !== null
    ) {
      message.localTimeoutStamp = object.localTimeoutStamp;
    } else {
      message.localTimeoutStamp = 0;
    }
    if (object.queryHeight !== undefined && object.queryHeight !== null) {
      message.queryHeight = object.queryHeight;
    } else {
      message.queryHeight = 0;
    }
    if (object.clientId !== undefined && object.clientId !== null) {
      message.clientId = object.clientId;
    } else {
      message.clientId = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.sourcePort !== undefined && object.sourcePort !== null) {
      message.sourcePort = object.sourcePort;
    } else {
      message.sourcePort = "";
    }
    if (object.sourceChannel !== undefined && object.sourceChannel !== null) {
      message.sourceChannel = object.sourceChannel;
    } else {
      message.sourceChannel = "";
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResponse: object = { queryId: "", capKey: 0 };

export const MsgSubmitCrossChainQueryResponse = {
  encode(
    message: MsgSubmitCrossChainQueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.queryId !== "") {
      writer.uint32(10).string(message.queryId);
    }
    if (message.capKey !== 0) {
      writer.uint32(16).uint64(message.capKey);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCrossChainQueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCrossChainQueryResponse,
    } as MsgSubmitCrossChainQueryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.queryId = reader.string();
          break;
        case 2:
          message.capKey = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitCrossChainQueryResponse {
    const message = {
      ...baseMsgSubmitCrossChainQueryResponse,
    } as MsgSubmitCrossChainQueryResponse;
    if (object.queryId !== undefined && object.queryId !== null) {
      message.queryId = String(object.queryId);
    } else {
      message.queryId = "";
    }
    if (object.capKey !== undefined && object.capKey !== null) {
      message.capKey = Number(object.capKey);
    } else {
      message.capKey = 0;
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQueryResponse): unknown {
    const obj: any = {};
    message.queryId !== undefined && (obj.queryId = message.queryId);
    message.capKey !== undefined && (obj.capKey = message.capKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQueryResponse>
  ): MsgSubmitCrossChainQueryResponse {
    const message = {
      ...baseMsgSubmitCrossChainQueryResponse,
    } as MsgSubmitCrossChainQueryResponse;
    if (object.queryId !== undefined && object.queryId !== null) {
      message.queryId = object.queryId;
    } else {
      message.queryId = "";
    }
    if (object.capKey !== undefined && object.capKey !== null) {
      message.capKey = object.capKey;
    } else {
      message.capKey = 0;
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResult: object = {
  id: "",
  path: "",
  queryHeight: 0,
  result: 0,
  sender: "",
};

export const MsgSubmitCrossChainQueryResult = {
  encode(
    message: MsgSubmitCrossChainQueryResult,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.path !== "") {
      writer.uint32(18).string(message.path);
    }
    if (message.queryHeight !== 0) {
      writer.uint32(24).uint64(message.queryHeight);
    }
    if (message.result !== 0) {
      writer.uint32(32).int32(message.result);
    }
    if (message.data.length !== 0) {
      writer.uint32(42).bytes(message.data);
    }
    if (message.sender !== "") {
      writer.uint32(50).string(message.sender);
    }
    for (const v of message.proofSpecs) {
      ProofSpec.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCrossChainQueryResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCrossChainQueryResult,
    } as MsgSubmitCrossChainQueryResult;
    message.proofSpecs = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.path = reader.string();
          break;
        case 3:
          message.queryHeight = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.result = reader.int32() as any;
          break;
        case 5:
          message.data = reader.bytes();
          break;
        case 6:
          message.sender = reader.string();
          break;
        case 7:
          message.proofSpecs.push(ProofSpec.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitCrossChainQueryResult {
    const message = {
      ...baseMsgSubmitCrossChainQueryResult,
    } as MsgSubmitCrossChainQueryResult;
    message.proofSpecs = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = String(object.path);
    } else {
      message.path = "";
    }
    if (object.queryHeight !== undefined && object.queryHeight !== null) {
      message.queryHeight = Number(object.queryHeight);
    } else {
      message.queryHeight = 0;
    }
    if (object.result !== undefined && object.result !== null) {
      message.result = queryResultFromJSON(object.result);
    } else {
      message.result = 0;
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = bytesFromBase64(object.data);
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    if (object.proofSpecs !== undefined && object.proofSpecs !== null) {
      for (const e of object.proofSpecs) {
        message.proofSpecs.push(ProofSpec.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQueryResult): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.path !== undefined && (obj.path = message.path);
    message.queryHeight !== undefined &&
      (obj.queryHeight = message.queryHeight);
    message.result !== undefined &&
      (obj.result = queryResultToJSON(message.result));
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    message.sender !== undefined && (obj.sender = message.sender);
    if (message.proofSpecs) {
      obj.proofSpecs = message.proofSpecs.map((e) =>
        e ? ProofSpec.toJSON(e) : undefined
      );
    } else {
      obj.proofSpecs = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQueryResult>
  ): MsgSubmitCrossChainQueryResult {
    const message = {
      ...baseMsgSubmitCrossChainQueryResult,
    } as MsgSubmitCrossChainQueryResult;
    message.proofSpecs = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.path !== undefined && object.path !== null) {
      message.path = object.path;
    } else {
      message.path = "";
    }
    if (object.queryHeight !== undefined && object.queryHeight !== null) {
      message.queryHeight = object.queryHeight;
    } else {
      message.queryHeight = 0;
    }
    if (object.result !== undefined && object.result !== null) {
      message.result = object.result;
    } else {
      message.result = 0;
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = new Uint8Array();
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    if (object.proofSpecs !== undefined && object.proofSpecs !== null) {
      for (const e of object.proofSpecs) {
        message.proofSpecs.push(ProofSpec.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResultResponse: object = {};

export const MsgSubmitCrossChainQueryResultResponse = {
  encode(
    _: MsgSubmitCrossChainQueryResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCrossChainQueryResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCrossChainQueryResultResponse,
    } as MsgSubmitCrossChainQueryResultResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitCrossChainQueryResultResponse {
    const message = {
      ...baseMsgSubmitCrossChainQueryResultResponse,
    } as MsgSubmitCrossChainQueryResultResponse;
    return message;
  },

  toJSON(_: MsgSubmitCrossChainQueryResultResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitCrossChainQueryResultResponse>
  ): MsgSubmitCrossChainQueryResultResponse {
    const message = {
      ...baseMsgSubmitCrossChainQueryResultResponse,
    } as MsgSubmitCrossChainQueryResultResponse;
    return message;
  },
};

/** Msg */
export interface Msg {
  /** submit query request */
  SubmitCrossChainQuery(
    request: MsgSubmitCrossChainQuery
  ): Promise<MsgSubmitCrossChainQueryResponse>;
  /** submit query result */
  SubmitCrossChainQueryResult(
    request: MsgSubmitCrossChainQueryResult
  ): Promise<MsgSubmitCrossChainQueryResultResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SubmitCrossChainQuery(
    request: MsgSubmitCrossChainQuery
  ): Promise<MsgSubmitCrossChainQueryResponse> {
    const data = MsgSubmitCrossChainQuery.encode(request).finish();
    const promise = this.rpc.request(
      "interchain_queries.ibc_query.v1.Msg",
      "SubmitCrossChainQuery",
      data
    );
    return promise.then((data) =>
      MsgSubmitCrossChainQueryResponse.decode(new Reader(data))
    );
  }

  SubmitCrossChainQueryResult(
    request: MsgSubmitCrossChainQueryResult
  ): Promise<MsgSubmitCrossChainQueryResultResponse> {
    const data = MsgSubmitCrossChainQueryResult.encode(request).finish();
    const promise = this.rpc.request(
      "interchain_queries.ibc_query.v1.Msg",
      "SubmitCrossChainQueryResult",
      data
    );
    return promise.then((data) =>
      MsgSubmitCrossChainQueryResultResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  for (let i = 0; i < arr.byteLength; ++i) {
    bin.push(String.fromCharCode(arr[i]));
  }
  return btoa(bin.join(""));
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
