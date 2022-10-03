/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

/**
 *  - NO_HASH: NO_HASH is the default if no data passed. Note this is an illegal argument some places.
 */
export enum Ics23HashOp {
  NO_HASH = "NO_HASH",
  SHA256 = "SHA256",
  SHA512 = "SHA512",
  KECCAK = "KECCAK",
  RIPEMD160 = "RIPEMD160",
  BITCOIN = "BITCOIN",
}

/**
* InnerSpec contains all store-specific structure info to determine if two proofs from a
given store are neighbors.

This enables:

isLeftMost(spec: InnerSpec, op: InnerOp)
isRightMost(spec: InnerSpec, op: InnerOp)
isLeftNeighbor(spec: InnerSpec, left: InnerOp, right: InnerOp)
*/
export interface Ics23InnerSpec {
  child_order?: number[];

  /** @format int32 */
  child_size?: number;

  /** @format int32 */
  min_prefix_length?: number;

  /** @format int32 */
  max_prefix_length?: number;

  /** @format byte */
  empty_child?: string;

  /**  - NO_HASH: NO_HASH is the default if no data passed. Note this is an illegal argument some places. */
  hash?: Ics23HashOp;
}

/**
* *
LeafOp represents the raw key-value data we wish to prove, and
must be flexible to represent the internal transformation from
the original key-value pairs into the basis hash, for many existing
merkle trees.

key and value are passed in. So that the signature of this operation is:
leafOp(key, value) -> output

To process this, first prehash the keys and values if needed (ANY means no hash in this case):
hkey = prehashKey(key)
hvalue = prehashValue(value)

Then combine the bytes, and hash it
output = hash(prefix || length(hkey) || hkey || length(hvalue) || hvalue)
*/
export interface Ics23LeafOp {
  /**  - NO_HASH: NO_HASH is the default if no data passed. Note this is an illegal argument some places. */
  hash?: Ics23HashOp;

  /**  - NO_HASH: NO_HASH is the default if no data passed. Note this is an illegal argument some places. */
  prehash_key?: Ics23HashOp;

  /**  - NO_HASH: NO_HASH is the default if no data passed. Note this is an illegal argument some places. */
  prehash_value?: Ics23HashOp;

  /**
   * - NO_PREFIX: NO_PREFIX don't include any length info
   *  - VAR_PROTO: VAR_PROTO uses protobuf (and go-amino) varint encoding of the length
   *  - VAR_RLP: VAR_RLP uses rlp int encoding of the length
   *  - FIXED32_BIG: FIXED32_BIG uses big-endian encoding of the length as a 32 bit integer
   *  - FIXED32_LITTLE: FIXED32_LITTLE uses little-endian encoding of the length as a 32 bit integer
   *  - FIXED64_BIG: FIXED64_BIG uses big-endian encoding of the length as a 64 bit integer
   *  - FIXED64_LITTLE: FIXED64_LITTLE uses little-endian encoding of the length as a 64 bit integer
   *  - REQUIRE_32_BYTES: REQUIRE_32_BYTES is like NONE, but will fail if the input is not exactly 32 bytes (sha256 output)
   *  - REQUIRE_64_BYTES: REQUIRE_64_BYTES is like NONE, but will fail if the input is not exactly 64 bytes (sha512 output)
   */
  length?: Ics23LengthOp;

  /**
   * prefix is a fixed bytes that may optionally be included at the beginning to differentiate
   * a leaf node from an inner node.
   * @format byte
   */
  prefix?: string;
}

/**
* - NO_PREFIX: NO_PREFIX don't include any length info
 - VAR_PROTO: VAR_PROTO uses protobuf (and go-amino) varint encoding of the length
 - VAR_RLP: VAR_RLP uses rlp int encoding of the length
 - FIXED32_BIG: FIXED32_BIG uses big-endian encoding of the length as a 32 bit integer
 - FIXED32_LITTLE: FIXED32_LITTLE uses little-endian encoding of the length as a 32 bit integer
 - FIXED64_BIG: FIXED64_BIG uses big-endian encoding of the length as a 64 bit integer
 - FIXED64_LITTLE: FIXED64_LITTLE uses little-endian encoding of the length as a 64 bit integer
 - REQUIRE_32_BYTES: REQUIRE_32_BYTES is like NONE, but will fail if the input is not exactly 32 bytes (sha256 output)
 - REQUIRE_64_BYTES: REQUIRE_64_BYTES is like NONE, but will fail if the input is not exactly 64 bytes (sha512 output)
*/
export enum Ics23LengthOp {
  NO_PREFIX = "NO_PREFIX",
  VAR_PROTO = "VAR_PROTO",
  VAR_RLP = "VAR_RLP",
  FIXED32BIG = "FIXED32_BIG",
  FIXED32LITTLE = "FIXED32_LITTLE",
  FIXED64BIG = "FIXED64_BIG",
  FIXED64LITTLE = "FIXED64_LITTLE",
  REQUIRE32BYTES = "REQUIRE_32_BYTES",
  REQUIRE64BYTES = "REQUIRE_64_BYTES",
}

/**
* *
ProofSpec defines what the expected parameters are for a given proof type.
This can be stored in the client and used to validate any incoming proofs.

verify(ProofSpec, Proof) -> Proof | Error

As demonstrated in tests, if we don't fix the algorithm used to calculate the
LeafHash for a given tree, there are many possible key-value pairs that can
generate a given hash (by interpretting the preimage differently).
We need this for proper security, requires client knows a priori what
tree format server uses. But not in code, rather a configuration object.
*/
export interface Ics23ProofSpec {
  /**
   * *
   * LeafOp represents the raw key-value data we wish to prove, and
   * must be flexible to represent the internal transformation from
   * the original key-value pairs into the basis hash, for many existing
   * merkle trees.
   *
   * key and value are passed in. So that the signature of this operation is:
   * leafOp(key, value) -> output
   *
   * To process this, first prehash the keys and values if needed (ANY means no hash in this case):
   * hkey = prehashKey(key)
   * hvalue = prehashValue(value)
   *
   * Then combine the bytes, and hash it
   * output = hash(prefix || length(hkey) || hkey || length(hvalue) || hvalue)
   */
  leaf_spec?: Ics23LeafOp;

  /**
   * InnerSpec contains all store-specific structure info to determine if two proofs from a
   * given store are neighbors.
   *
   * This enables:
   *
   * isLeftMost(spec: InnerSpec, op: InnerOp)
   * isRightMost(spec: InnerSpec, op: InnerOp)
   * isLeftNeighbor(spec: InnerSpec, left: InnerOp, right: InnerOp)
   */
  inner_spec?: Ics23InnerSpec;

  /** @format int32 */
  max_depth?: number;

  /** @format int32 */
  min_depth?: number;
}

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface ProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   *
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   *
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   *
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* Normally the RevisionHeight is incremented at each height while keeping
RevisionNumber the same. However some consensus algorithms may choose to
reset the height in certain conditions e.g. hard forks, state-machine
breaking changes In these cases, the RevisionNumber is incremented so that
height continues to be monitonically increasing even as the RevisionHeight
gets reset
*/
export interface V1Height {
  /** @format uint64 */
  revision_number?: string;

  /** @format uint64 */
  revision_height?: string;
}

export interface V1MsgSubmitCrossChainQueryResponse {
  query_id?: string;

  /** @format uint64 */
  cap_key?: string;
}

export type V1MsgSubmitCrossChainQueryResultResponse = object;

export interface V1QueryCrossChainQueryResultResponse {
  id?: string;

  /**
   * - QUERY_RESULT_UNSPECIFIED: UNSPECIFIED
   *  - QUERY_RESULT_SUCCESS: SUCCESS
   *  - QUERY_RESULT_FAILURE: FAILURE
   *  - QUERY_RESULT_TIMEOUT: TIMEOUT
   */
  result?: V1QueryResult;

  /** @format byte */
  data?: string;
}

/**
* - QUERY_RESULT_UNSPECIFIED: UNSPECIFIED
 - QUERY_RESULT_SUCCESS: SUCCESS
 - QUERY_RESULT_FAILURE: FAILURE
 - QUERY_RESULT_TIMEOUT: TIMEOUT
*/
export enum V1QueryResult {
  QUERY_RESULT_UNSPECIFIED = "QUERY_RESULT_UNSPECIFIED",
  QUERY_RESULT_SUCCESS = "QUERY_RESULT_SUCCESS",
  QUERY_RESULT_FAILURE = "QUERY_RESULT_FAILURE",
  QUERY_RESULT_TIMEOUT = "QUERY_RESULT_TIMEOUT",
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title ibc_query/v1/crosschainquery.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCrossChainQueryResult
   * @summary query CrossChainQueryResult
   * @request GET:/ibc/apps/ibc_query/v1/{id}
   */
  queryCrossChainQueryResult = (id: string, params: RequestParams = {}) =>
    this.request<V1QueryCrossChainQueryResultResponse, RpcStatus>({
      path: `/ibc/apps/ibc-query/v1/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
