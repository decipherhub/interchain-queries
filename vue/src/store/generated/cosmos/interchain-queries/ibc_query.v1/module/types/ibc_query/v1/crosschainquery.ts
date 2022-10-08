/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Height } from "../../ibc/core/client/v1/client";

export const protobufPackage = "ibc_query.v1";

/** QueryResult */
export enum QueryResult {
  /** QUERY_RESULT_UNSPECIFIED - UNSPECIFIED */
  QUERY_RESULT_UNSPECIFIED = 0,
  /** QUERY_RESULT_SUCCESS - SUCCESS */
  QUERY_RESULT_SUCCESS = 1,
  /** QUERY_RESULT_FAILURE - FAILURE */
  QUERY_RESULT_FAILURE = 2,
  /** QUERY_RESULT_TIMEOUT - TIMEOUT */
  QUERY_RESULT_TIMEOUT = 3,
  UNRECOGNIZED = -1,
}

export function queryResultFromJSON(object: any): QueryResult {
  switch (object) {
    case 0:
    case "QUERY_RESULT_UNSPECIFIED":
      return QueryResult.QUERY_RESULT_UNSPECIFIED;
    case 1:
    case "QUERY_RESULT_SUCCESS":
      return QueryResult.QUERY_RESULT_SUCCESS;
    case 2:
    case "QUERY_RESULT_FAILURE":
      return QueryResult.QUERY_RESULT_FAILURE;
    case 3:
    case "QUERY_RESULT_TIMEOUT":
      return QueryResult.QUERY_RESULT_TIMEOUT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return QueryResult.UNRECOGNIZED;
  }
}

export function queryResultToJSON(object: QueryResult): string {
  switch (object) {
    case QueryResult.QUERY_RESULT_UNSPECIFIED:
      return "QUERY_RESULT_UNSPECIFIED";
    case QueryResult.QUERY_RESULT_SUCCESS:
      return "QUERY_RESULT_SUCCESS";
    case QueryResult.QUERY_RESULT_FAILURE:
      return "QUERY_RESULT_FAILURE";
    case QueryResult.QUERY_RESULT_TIMEOUT:
      return "QUERY_RESULT_TIMEOUT";
    default:
      return "UNKNOWN";
  }
}

/** CrossChainQuery */
export interface CrossChainQuery {
  id: string;
  path: string;
  local_timeout_height: Height | undefined;
  local_timeout_timestamp: number;
  query_height: number;
  chain_id: string;
}

/** CrossChainQueryResult */
export interface CrossChainQueryResult {
  id: string;
  result: QueryResult;
  data: Uint8Array;
}

const baseCrossChainQuery: object = {
  id: "",
  path: "",
  local_timeout_timestamp: 0,
  query_height: 0,
  chain_id: "",
};

export const CrossChainQuery = {
  encode(message: CrossChainQuery, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.path !== "") {
      writer.uint32(18).string(message.path);
    }
    if (message.local_timeout_height !== undefined) {
      Height.encode(
        message.local_timeout_height,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.local_timeout_timestamp !== 0) {
      writer.uint32(32).uint64(message.local_timeout_timestamp);
    }
    if (message.query_height !== 0) {
      writer.uint32(40).uint64(message.query_height);
    }
    if (message.chain_id !== "") {
      writer.uint32(50).string(message.chain_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CrossChainQuery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCrossChainQuery } as CrossChainQuery;
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
          message.local_timeout_height = Height.decode(reader, reader.uint32());
          break;
        case 4:
          message.local_timeout_timestamp = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 5:
          message.query_height = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.chain_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CrossChainQuery {
    const message = { ...baseCrossChainQuery } as CrossChainQuery;
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
      object.local_timeout_timestamp !== undefined &&
      object.local_timeout_timestamp !== null
    ) {
      message.local_timeout_timestamp = Number(object.local_timeout_timestamp);
    } else {
      message.local_timeout_timestamp = 0;
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
    return message;
  },

  toJSON(message: CrossChainQuery): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.path !== undefined && (obj.path = message.path);
    message.local_timeout_height !== undefined &&
      (obj.local_timeout_height = message.local_timeout_height
        ? Height.toJSON(message.local_timeout_height)
        : undefined);
    message.local_timeout_timestamp !== undefined &&
      (obj.local_timeout_timestamp = message.local_timeout_timestamp);
    message.query_height !== undefined &&
      (obj.query_height = message.query_height);
    message.chain_id !== undefined && (obj.chain_id = message.chain_id);
    return obj;
  },

  fromPartial(object: DeepPartial<CrossChainQuery>): CrossChainQuery {
    const message = { ...baseCrossChainQuery } as CrossChainQuery;
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
      object.local_timeout_timestamp !== undefined &&
      object.local_timeout_timestamp !== null
    ) {
      message.local_timeout_timestamp = object.local_timeout_timestamp;
    } else {
      message.local_timeout_timestamp = 0;
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
    return message;
  },
};

const baseCrossChainQueryResult: object = { id: "", result: 0 };

export const CrossChainQueryResult = {
  encode(
    message: CrossChainQueryResult,
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

  decode(input: Reader | Uint8Array, length?: number): CrossChainQueryResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCrossChainQueryResult } as CrossChainQueryResult;
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

  fromJSON(object: any): CrossChainQueryResult {
    const message = { ...baseCrossChainQueryResult } as CrossChainQueryResult;
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

  toJSON(message: CrossChainQueryResult): unknown {
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
    object: DeepPartial<CrossChainQueryResult>
  ): CrossChainQueryResult {
    const message = { ...baseCrossChainQueryResult } as CrossChainQueryResult;
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
