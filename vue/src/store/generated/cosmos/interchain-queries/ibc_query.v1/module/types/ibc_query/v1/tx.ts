/* eslint-disable */
import {
  QueryResult,
  queryResultFromJSON,
  queryResultToJSON,
} from "../../ibc_query/v1/crosschainquery";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Height } from "../../ibc/core/client/v1/client";
import { ProofSpec } from "../../proofs";

export const protobufPackage = "ibc_query.v1";

/** MsgSubmitCrossChainQuery */
export interface MsgSubmitCrossChainQuery {
  path: string;
  local_timeout_height: Height | undefined;
  local_timeout_stamp: number;
  query_height: number;
  chain_id: string;
  sender: string;
}

/** MsgSubmitCrossChainQueryResponse */
export interface MsgSubmitCrossChainQueryResponse {
  query_id: string;
  cap_key: number;
}

/** MsgSubmitCrossChainQueryResult */
export interface MsgSubmitCrossChainQueryResult {
  id: string;
  path: string;
  query_height: number;
  result: QueryResult;
  data: Uint8Array;
  /** TODO: Proof specifications used in verifying counterparty state */
  proof_specs: ProofSpec[];
}

/** MsgSubmitCrossChainQueryResultResponse */
export interface MsgSubmitCrossChainQueryResultResponse {}

const baseMsgSubmitCrossChainQuery: object = {
  path: "",
  local_timeout_stamp: 0,
  query_height: 0,
  chain_id: "",
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
    if (message.local_timeout_height !== undefined) {
      Height.encode(
        message.local_timeout_height,
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.local_timeout_stamp !== 0) {
      writer.uint32(24).uint64(message.local_timeout_stamp);
    }
    if (message.query_height !== 0) {
      writer.uint32(32).uint64(message.query_height);
    }
    if (message.chain_id !== "") {
      writer.uint32(42).string(message.chain_id);
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
          message.local_timeout_height = Height.decode(reader, reader.uint32());
          break;
        case 3:
          message.local_timeout_stamp = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.query_height = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.chain_id = reader.string();
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
      object.local_timeout_height !== undefined &&
      object.local_timeout_height !== null
    ) {
      message.local_timeout_height = Height.fromJSON(
        object.local_timeout_height
      );
    } else {
      message.local_timeout_height = undefined;
    }
    if (
      object.local_timeout_stamp !== undefined &&
      object.local_timeout_stamp !== null
    ) {
      message.local_timeout_stamp = Number(object.local_timeout_stamp);
    } else {
      message.local_timeout_stamp = 0;
    }
    if (object.query_height !== undefined && object.query_height !== null) {
      message.query_height = Number(object.query_height);
    } else {
      message.query_height = 0;
    }
    if (object.chain_id !== undefined && object.chain_id !== null) {
      message.chain_id = String(object.chain_id);
    } else {
      message.chain_id = "";
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
    message.local_timeout_height !== undefined &&
      (obj.local_timeout_height = message.local_timeout_height
        ? Height.toJSON(message.local_timeout_height)
        : undefined);
    message.local_timeout_stamp !== undefined &&
      (obj.local_timeout_stamp = message.local_timeout_stamp);
    message.query_height !== undefined &&
      (obj.query_height = message.query_height);
    message.chain_id !== undefined && (obj.chain_id = message.chain_id);
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
      object.local_timeout_height !== undefined &&
      object.local_timeout_height !== null
    ) {
      message.local_timeout_height = Height.fromPartial(
        object.local_timeout_height
      );
    } else {
      message.local_timeout_height = undefined;
    }
    if (
      object.local_timeout_stamp !== undefined &&
      object.local_timeout_stamp !== null
    ) {
      message.local_timeout_stamp = object.local_timeout_stamp;
    } else {
      message.local_timeout_stamp = 0;
    }
    if (object.query_height !== undefined && object.query_height !== null) {
      message.query_height = object.query_height;
    } else {
      message.query_height = 0;
    }
    if (object.chain_id !== undefined && object.chain_id !== null) {
      message.chain_id = object.chain_id;
    } else {
      message.chain_id = "";
    }
    if (object.sender !== undefined && object.sender !== null) {
      message.sender = object.sender;
    } else {
      message.sender = "";
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResponse: object = {
  query_id: "",
  cap_key: 0,
};

export const MsgSubmitCrossChainQueryResponse = {
  encode(
    message: MsgSubmitCrossChainQueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.query_id !== "") {
      writer.uint32(10).string(message.query_id);
    }
    if (message.cap_key !== 0) {
      writer.uint32(16).uint64(message.cap_key);
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
          message.query_id = reader.string();
          break;
        case 2:
          message.cap_key = longToNumber(reader.uint64() as Long);
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
    if (object.query_id !== undefined && object.query_id !== null) {
      message.query_id = String(object.query_id);
    } else {
      message.query_id = "";
    }
    if (object.cap_key !== undefined && object.cap_key !== null) {
      message.cap_key = Number(object.cap_key);
    } else {
      message.cap_key = 0;
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQueryResponse): unknown {
    const obj: any = {};
    message.query_id !== undefined && (obj.query_id = message.query_id);
    message.cap_key !== undefined && (obj.cap_key = message.cap_key);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQueryResponse>
  ): MsgSubmitCrossChainQueryResponse {
    const message = {
      ...baseMsgSubmitCrossChainQueryResponse,
    } as MsgSubmitCrossChainQueryResponse;
    if (object.query_id !== undefined && object.query_id !== null) {
      message.query_id = object.query_id;
    } else {
      message.query_id = "";
    }
    if (object.cap_key !== undefined && object.cap_key !== null) {
      message.cap_key = object.cap_key;
    } else {
      message.cap_key = 0;
    }
    return message;
  },
};

const baseMsgSubmitCrossChainQueryResult: object = {
  id: "",
  path: "",
  query_height: 0,
  result: 0,
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
    if (message.query_height !== 0) {
      writer.uint32(24).uint64(message.query_height);
    }
    if (message.result !== 0) {
      writer.uint32(32).int32(message.result);
    }
    if (message.data.length !== 0) {
      writer.uint32(42).bytes(message.data);
    }
    for (const v of message.proof_specs) {
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
    message.proof_specs = [];
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
          message.query_height = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.result = reader.int32() as any;
          break;
        case 5:
          message.data = reader.bytes();
          break;
        case 6:
          message.proof_specs.push(ProofSpec.decode(reader, reader.uint32()));
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
    message.proof_specs = [];
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
    if (object.query_height !== undefined && object.query_height !== null) {
      message.query_height = Number(object.query_height);
    } else {
      message.query_height = 0;
    }
    if (object.result !== undefined && object.result !== null) {
      message.result = queryResultFromJSON(object.result);
    } else {
      message.result = 0;
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = bytesFromBase64(object.data);
    }
    if (object.proof_specs !== undefined && object.proof_specs !== null) {
      for (const e of object.proof_specs) {
        message.proof_specs.push(ProofSpec.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgSubmitCrossChainQueryResult): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.path !== undefined && (obj.path = message.path);
    message.query_height !== undefined &&
      (obj.query_height = message.query_height);
    message.result !== undefined &&
      (obj.result = queryResultToJSON(message.result));
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    if (message.proof_specs) {
      obj.proof_specs = message.proof_specs.map((e) =>
        e ? ProofSpec.toJSON(e) : undefined
      );
    } else {
      obj.proof_specs = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCrossChainQueryResult>
  ): MsgSubmitCrossChainQueryResult {
    const message = {
      ...baseMsgSubmitCrossChainQueryResult,
    } as MsgSubmitCrossChainQueryResult;
    message.proof_specs = [];
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
    if (object.query_height !== undefined && object.query_height !== null) {
      message.query_height = object.query_height;
    } else {
      message.query_height = 0;
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
    if (object.proof_specs !== undefined && object.proof_specs !== null) {
      for (const e of object.proof_specs) {
        message.proof_specs.push(ProofSpec.fromPartial(e));
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
      "ibc_query.v1.Msg",
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
      "ibc_query.v1.Msg",
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
