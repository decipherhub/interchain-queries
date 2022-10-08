/* eslint-disable */
import {
  CrossChainQuery,
  CrossChainQueryResult,
} from "../../ibc_query/v1/crosschainquery";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "ibc.applications.ibc_query.v1";

/** GenesisState defines the ICS31 ibc-query genesis state */
export interface GenesisState {
  queries: CrossChainQuery[];
  results: CrossChainQueryResult[];
  port_id: string;
}

const baseGenesisState: object = { port_id: "" };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.queries) {
      CrossChainQuery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.results) {
      CrossChainQueryResult.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.port_id !== "") {
      writer.uint32(26).string(message.port_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.queries = [];
    message.results = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.queries.push(CrossChainQuery.decode(reader, reader.uint32()));
          break;
        case 2:
          message.results.push(
            CrossChainQueryResult.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.port_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.queries = [];
    message.results = [];
    if (object.queries !== undefined && object.queries !== null) {
      for (const e of object.queries) {
        message.queries.push(CrossChainQuery.fromJSON(e));
      }
    }
    if (object.results !== undefined && object.results !== null) {
      for (const e of object.results) {
        message.results.push(CrossChainQueryResult.fromJSON(e));
      }
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = String(object.port_id);
    } else {
      message.port_id = "";
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.queries) {
      obj.queries = message.queries.map((e) =>
        e ? CrossChainQuery.toJSON(e) : undefined
      );
    } else {
      obj.queries = [];
    }
    if (message.results) {
      obj.results = message.results.map((e) =>
        e ? CrossChainQueryResult.toJSON(e) : undefined
      );
    } else {
      obj.results = [];
    }
    message.port_id !== undefined && (obj.port_id = message.port_id);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.queries = [];
    message.results = [];
    if (object.queries !== undefined && object.queries !== null) {
      for (const e of object.queries) {
        message.queries.push(CrossChainQuery.fromPartial(e));
      }
    }
    if (object.results !== undefined && object.results !== null) {
      for (const e of object.results) {
        message.results.push(CrossChainQueryResult.fromPartial(e));
      }
    }
    if (object.port_id !== undefined && object.port_id !== null) {
      message.port_id = object.port_id;
    } else {
      message.port_id = "";
    }
    return message;
  },
};

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
