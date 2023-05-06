// package: items
// file: items.proto

import * as items_pb from "./items_pb";
import grpc from "grpc-web";

type ItemsServiceCreate = {
  readonly methodName: string;
  readonly service: typeof ItemsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof items_pb.Item;
  readonly responseType: typeof items_pb.Item;
};

type ItemsServiceGet = {
  readonly methodName: string;
  readonly service: typeof ItemsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof items_pb.Item;
  readonly responseType: typeof items_pb.Item;
};

type ItemsServiceUpdate = {
  readonly methodName: string;
  readonly service: typeof ItemsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof items_pb.Item;
  readonly responseType: typeof items_pb.Item;
};

type ItemsServiceDelete = {
  readonly methodName: string;
  readonly service: typeof ItemsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof items_pb.Item;
  readonly responseType: typeof items_pb.Item;
};

type ItemsServiceGetAll = {
  readonly methodName: string;
  readonly service: typeof ItemsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof items_pb.Empty;
  readonly responseType: typeof items_pb.ItemsList;
};

export class ItemsService {
  static readonly serviceName: string;
  static readonly Create: ItemsServiceCreate;
  static readonly Get: ItemsServiceGet;
  static readonly Update: ItemsServiceUpdate;
  static readonly Delete: ItemsServiceDelete;
  static readonly GetAll: ItemsServiceGetAll;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ItemsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  create(
    requestMessage: items_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  create(
    requestMessage: items_pb.Item,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  get(
    requestMessage: items_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  get(
    requestMessage: items_pb.Item,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  update(
    requestMessage: items_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  update(
    requestMessage: items_pb.Item,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: items_pb.Item,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  delete(
    requestMessage: items_pb.Item,
    callback: (error: ServiceError|null, responseMessage: items_pb.Item|null) => void
  ): UnaryResponse;
  getAll(
    requestMessage: items_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: items_pb.ItemsList|null) => void
  ): UnaryResponse;
  getAll(
    requestMessage: items_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: items_pb.ItemsList|null) => void
  ): UnaryResponse;
}

