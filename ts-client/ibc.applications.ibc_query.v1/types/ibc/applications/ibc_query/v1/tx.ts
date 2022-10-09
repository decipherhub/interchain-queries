/* eslint-disable */
import {
  QueryResult,
  queryResultFromJSON,
  queryResultToJSON,
} from "../../../../ibc/applications/ibc_query/v1/crosschainquery";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Height } from "../../../../ibc/core/client/v1/client";
import { Capability } from "../../../../cosmos/capability/v1beta1/capability";
import { ProofSpec } from "../../../../proofs";

export const protobufPackage = "ibc.applications.ibc_query.v1";

/** MsgSubmitCrossChainQuery */
export interface MsgSubmitCrossChainQuery {
  path: string;
  localTimeoutHeight: Height | undefined;
  localTimeoutStamp: number;
  queryHeight: number;
  chainId: string;
  sender: string;
}

/** MsgSubmitCrossChainQueryResponse */
export interface MsgSubmitCrossChainQueryResponse {
  id: string;
  capKey: Capability | undefined;
}

/** MsgSubmitCrossChainQueryResult */
export interface MsgSubmitCrossChainQueryResult {
  id: string;
  queryHeight: number;
  result: QueryResult;
  data: Uint8Array;
  sender: string;
  /** TODO: Proof specifications used in verifying counterparty state */
  proofSpecs: ProofSpec[];
}

/** MsgSubmitCrossChainQueryResultResponse */
export interface MsgSubmitCrossChainQueryResultResponse {}

/** MsgSubmitPruneCrossChainQueryResult */
export interface MsgSubmitPruneCrossChainQueryResult {
  id: string;
  capKey: Capability | undefined;
  sender: string;
}

/** MsgSubmitPruneCrossChainQueryResultResponse */
export interface MsgSubmitPruneCrossChainQueryResultResponse {
  id: string;
  result: QueryResult;
  data: Uint8Array;
}

const baseMsgSubmitCrossChainQuery: object = {
  path: "",
  localTimeoutStamp: 0,
  queryHeight: 0,
  chainId: "",
  sender: "",
};

export const MsgSubmitCrossChainQuery = {
  encode(
    message: MsgSubmitCrossChainQuery,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.path !== "") {
      writer.uint32(10).string(message.path);
    }
    if (message.localTimeoutHeight !== undefined) {
      Height.encode(
        message.localTimeoutHeight,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.localTimeoutStamp !== 0) {
      writer.uint32(24).uint64(message.localTimeoutStamp);
    }
    if (message.queryHeight !== 0) {
      writer.uint32(32).uint64(message.queryHeight);
    }
    if (message.chainId !== "") {
      writer.uint32(42).string(message.chainId);
    }
    if (message.sender !== "") {
      writer.uint32(50).string(message.sender);
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
          message.path = reader.string();
          break;
        case 2:
          message.localTimeoutHeight = Height.decode(reader, reader.uint32());
          break;
        case 3:
          message.localTimeoutStamp = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.queryHeight = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.chainId = reader.string();
          break;
        case 6:
          message.sender = reader.string();
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
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = String(object.chainId);
    } else {
      message.chainId = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQuery): unknown {
    const obj: any = {};
    message.path !== undefined && (obj.path = message.path);
    message.localTimeoutHeight !== undefined &&
      (obj.localTimeoutHeight = message.localTimeoutHeight
        ? Height.toJSON(message.localTimeoutHeight)
        : undefined);
    message.localTimeoutStamp !== undefined &&
      (obj.localTimeoutStamp = message.localTimeoutStamp);
    message.queryHeight !== undefined &&
      (obj.queryHeight = message.queryHeight);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQuery>
  ): MsgSubmitCrossChainQuery {
    const message = {
      ...baseMsgSubmitCrossChainQuery,
    } as MsgSubmitCrossChainQuery;
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
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResponse: object = { id: "" };

export const MsgSubmitCrossChainQueryResponse = {
  encode(
    message: MsgSubmitCrossChainQueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.capKey !== undefined) {
      Capability.encode(message.capKey, writer.uint32(18).fork()).ldelim();
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
          message.id = reader.string();
          break;
        case 2:
          message.capKey = Capability.decode(reader, reader.uint32());
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
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.capKey !== undefined && object.capKey !== null) {
      message.capKey = Capability.fromJSON(object.capKey);
    } else {
      message.capKey = undefined;
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQueryResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.capKey !== undefined &&
      (obj.capKey = message.capKey
        ? Capability.toJSON(message.capKey)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQueryResponse>
  ): MsgSubmitCrossChainQueryResponse {
    const message = {
      ...baseMsgSubmitCrossChainQueryResponse,
    } as MsgSubmitCrossChainQueryResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.capKey !== undefined && object.capKey !== null) {
      message.capKey = Capability.fromPartial(object.capKey);
    } else {
      message.capKey = undefined;
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResult: object = {
  id: "",
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
    if (message.queryHeight !== 0) {
      writer.uint32(16).uint64(message.queryHeight);
    }
    if (message.result !== 0) {
      writer.uint32(24).int32(message.result);
    }
    if (message.data.length !== 0) {
      writer.uint32(34).bytes(message.data);
    }
    if (message.sender !== "") {
      writer.uint32(42).string(message.sender);
    }
    for (const v of message.proofSpecs) {
      ProofSpec.encode(v!, writer.uint32(50).fork()).ldelim();
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
          message.queryHeight = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.result = reader.int32() as any;
          break;
        case 4:
          message.data = reader.bytes();
          break;
        case 5:
          message.sender = reader.string();
          break;
        case 6:
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

const baseMsgSubmitPruneCrossChainQueryResult: object = { id: "", sender: "" };

export const MsgSubmitPruneCrossChainQueryResult = {
  encode(
    message: MsgSubmitPruneCrossChainQueryResult,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.capKey !== undefined) {
      Capability.encode(message.capKey, writer.uint32(18).fork()).ldelim();
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitPruneCrossChainQueryResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitPruneCrossChainQueryResult,
    } as MsgSubmitPruneCrossChainQueryResult;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.capKey = Capability.decode(reader, reader.uint32());
          break;
        case 3:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitPruneCrossChainQueryResult {
    const message = {
      ...baseMsgSubmitPruneCrossChainQueryResult,
    } as MsgSubmitPruneCrossChainQueryResult;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.capKey !== undefined && object.capKey !== null) {
      message.capKey = Capability.fromJSON(object.capKey);
    } else {
      message.capKey = undefined;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = String(object.sender);
    } else {
      message.sender = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitPruneCrossChainQueryResult): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.capKey !== undefined &&
      (obj.capKey = message.capKey
        ? Capability.toJSON(message.capKey)
        : undefined);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitPruneCrossChainQueryResult>
  ): MsgSubmitPruneCrossChainQueryResult {
    const message = {
      ...baseMsgSubmitPruneCrossChainQueryResult,
    } as MsgSubmitPruneCrossChainQueryResult;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.capKey !== undefined && object.capKey !== null) {
      message.capKey = Capability.fromPartial(object.capKey);
    } else {
      message.capKey = undefined;
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseMsgSubmitPruneCrossChainQueryResultResponse: object = {
  id: "",
  result: 0,
};

export const MsgSubmitPruneCrossChainQueryResultResponse = {
  encode(
    message: MsgSubmitPruneCrossChainQueryResultResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.result !== 0) {
      writer.uint32(16).int32(message.result);
    }
    if (message.data.length !== 0) {
      writer.uint32(26).bytes(message.data);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitPruneCrossChainQueryResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitPruneCrossChainQueryResultResponse,
    } as MsgSubmitPruneCrossChainQueryResultResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.result = reader.int32() as any;
          break;
        case 3:
          message.data = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitPruneCrossChainQueryResultResponse {
    const message = {
      ...baseMsgSubmitPruneCrossChainQueryResultResponse,
    } as MsgSubmitPruneCrossChainQueryResultResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.result !== undefined && object.result !== null) {
      message.result = queryResultFromJSON(object.result);
    } else {
      message.result = 0;
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = bytesFromBase64(object.data);
    }
    return message;
  },

  toJSON(message: MsgSubmitPruneCrossChainQueryResultResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.result !== undefined &&
      (obj.result = queryResultToJSON(message.result));
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitPruneCrossChainQueryResultResponse>
  ): MsgSubmitPruneCrossChainQueryResultResponse {
    const message = {
      ...baseMsgSubmitPruneCrossChainQueryResultResponse,
    } as MsgSubmitPruneCrossChainQueryResultResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
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
  /** retrieve the result and delete it */
  SubmitPruneCrossChainQueryResult(
    request: MsgSubmitPruneCrossChainQueryResult
  ): Promise<MsgSubmitPruneCrossChainQueryResultResponse>;
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
      "ibc.applications.ibc_query.v1.Msg",
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
      "ibc.applications.ibc_query.v1.Msg",
      "SubmitCrossChainQueryResult",
      data
    );
    return promise.then((data) =>
      MsgSubmitCrossChainQueryResultResponse.decode(new Reader(data))
    );
  }

  SubmitPruneCrossChainQueryResult(
    request: MsgSubmitPruneCrossChainQueryResult
  ): Promise<MsgSubmitPruneCrossChainQueryResultResponse> {
    const data = MsgSubmitPruneCrossChainQueryResult.encode(request).finish();
    const promise = this.rpc.request(
      "ibc.applications.ibc_query.v1.Msg",
      "SubmitPruneCrossChainQueryResult",
      data
    );
    return promise.then((data) =>
      MsgSubmitPruneCrossChainQueryResultResponse.decode(new Reader(data))
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
