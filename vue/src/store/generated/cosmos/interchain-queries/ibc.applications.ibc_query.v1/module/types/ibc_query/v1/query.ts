/* eslint-disable */
import {
  QueryResult,
  CrossChainQuery,
  queryResultFromJSON,
  queryResultToJSON,
} from "../../ibc_query/v1/crosschainquery";
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "ibc.applications.ibc_query.v1";

/** QueryCrossChainQuery */
export interface QueryCrossChainQuery {
  /** query id */
  id: string;
}

/** QueryCrossChainQueryResponse */
export interface QueryCrossChainQueryResponse {
  result: CrossChainQuery | undefined;
}

/** QueryCrossChainQueryResult */
export interface QueryCrossChainQueryResult {
  /** query id */
  id: string;
}

/** QueryCrossChainQueryResultResponse */
export interface QueryCrossChainQueryResultResponse {
  id: string;
  result: QueryResult;
  data: Uint8Array;
}

const baseQueryCrossChainQuery: object = { id: "" };

export const QueryCrossChainQuery = {
  encode(
    message: QueryCrossChainQuery,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryCrossChainQuery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryCrossChainQuery } as QueryCrossChainQuery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCrossChainQuery {
    const message = { ...baseQueryCrossChainQuery } as QueryCrossChainQuery;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: QueryCrossChainQuery): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryCrossChainQuery>): QueryCrossChainQuery {
    const message = { ...baseQueryCrossChainQuery } as QueryCrossChainQuery;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseQueryCrossChainQueryResponse: object = {};

export const QueryCrossChainQueryResponse = {
  encode(
    message: QueryCrossChainQueryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.result !== undefined) {
      CrossChainQuery.encode(message.result, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCrossChainQueryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCrossChainQueryResponse,
    } as QueryCrossChainQueryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.result = CrossChainQuery.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCrossChainQueryResponse {
    const message = {
      ...baseQueryCrossChainQueryResponse,
    } as QueryCrossChainQueryResponse;
    if (object.result !== undefined && object.result !== null) {
      message.result = CrossChainQuery.fromJSON(object.result);
    } else {
      message.result = undefined;
    }
    return message;
  },

  toJSON(message: QueryCrossChainQueryResponse): unknown {
    const obj: any = {};
    message.result !== undefined &&
      (obj.result = message.result
        ? CrossChainQuery.toJSON(message.result)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCrossChainQueryResponse>
  ): QueryCrossChainQueryResponse {
    const message = {
      ...baseQueryCrossChainQueryResponse,
    } as QueryCrossChainQueryResponse;
    if (object.result !== undefined && object.result !== null) {
      message.result = CrossChainQuery.fromPartial(object.result);
    } else {
      message.result = undefined;
    }
    return message;
  },
};

const baseQueryCrossChainQueryResult: object = { id: "" };

export const QueryCrossChainQueryResult = {
  encode(
    message: QueryCrossChainQueryResult,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCrossChainQueryResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCrossChainQueryResult,
    } as QueryCrossChainQueryResult;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCrossChainQueryResult {
    const message = {
      ...baseQueryCrossChainQueryResult,
    } as QueryCrossChainQueryResult;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    return message;
  },

  toJSON(message: QueryCrossChainQueryResult): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCrossChainQueryResult>
  ): QueryCrossChainQueryResult {
    const message = {
      ...baseQueryCrossChainQueryResult,
    } as QueryCrossChainQueryResult;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    return message;
  },
};

const baseQueryCrossChainQueryResultResponse: object = { id: "", result: 0 };

export const QueryCrossChainQueryResultResponse = {
  encode(
    message: QueryCrossChainQueryResultResponse,
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
  ): QueryCrossChainQueryResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCrossChainQueryResultResponse,
    } as QueryCrossChainQueryResultResponse;
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

  fromJSON(object: any): QueryCrossChainQueryResultResponse {
    const message = {
      ...baseQueryCrossChainQueryResultResponse,
    } as QueryCrossChainQueryResultResponse;
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

  toJSON(message: QueryCrossChainQueryResultResponse): unknown {
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
    object: DeepPartial<QueryCrossChainQueryResultResponse>
  ): QueryCrossChainQueryResultResponse {
    const message = {
      ...baseQueryCrossChainQueryResultResponse,
    } as QueryCrossChainQueryResultResponse;
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

/** Query */
export interface Query {
  /** query CrossChainQuery */
  CrossChainQuery(
    request: QueryCrossChainQuery
  ): Promise<QueryCrossChainQueryResponse>;
  /** query CrossChainQueryResult */
  CrossChainQueryResult(
    request: QueryCrossChainQueryResult
  ): Promise<QueryCrossChainQueryResultResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CrossChainQuery(
    request: QueryCrossChainQuery
  ): Promise<QueryCrossChainQueryResponse> {
    const data = QueryCrossChainQuery.encode(request).finish();
    const promise = this.rpc.request(
      "ibc.applications.ibc_query.v1.Query",
      "CrossChainQuery",
      data
    );
    return promise.then((data) =>
      QueryCrossChainQueryResponse.decode(new Reader(data))
    );
  }

  CrossChainQueryResult(
    request: QueryCrossChainQueryResult
  ): Promise<QueryCrossChainQueryResultResponse> {
    const data = QueryCrossChainQueryResult.encode(request).finish();
    const promise = this.rpc.request(
      "ibc.applications.ibc_query.v1.Query",
      "CrossChainQueryResult",
      data
    );
    return promise.then((data) =>
      QueryCrossChainQueryResultResponse.decode(new Reader(data))
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
