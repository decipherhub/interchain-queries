/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "interchain_queries.ibc_query.v1";

/** IBCQueryPacketData defines a struct for the cross chain query packet payload */
export interface IBCQueryPacketData {
  id: string;
  path: string;
  queryHeight: number;
}

const baseIBCQueryPacketData: object = { id: "", path: "", queryHeight: 0 };

export const IBCQueryPacketData = {
  encode(
    message: IBCQueryPacketData,
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
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): IBCQueryPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseIBCQueryPacketData } as IBCQueryPacketData;
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IBCQueryPacketData {
    const message = { ...baseIBCQueryPacketData } as IBCQueryPacketData;
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
    return message;
  },

  toJSON(message: IBCQueryPacketData): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.path !== undefined && (obj.path = message.path);
    message.queryHeight !== undefined &&
      (obj.queryHeight = message.queryHeight);
    return obj;
  },

  fromPartial(object: DeepPartial<IBCQueryPacketData>): IBCQueryPacketData {
    const message = { ...baseIBCQueryPacketData } as IBCQueryPacketData;
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
