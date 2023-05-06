// Original file: ../items.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Empty as _items_Empty, Empty__Output as _items_Empty__Output } from '../items/Empty';
import type { Item as _items_Item, Item__Output as _items_Item__Output } from '../items/Item';
import type { ItemsList as _items_ItemsList, ItemsList__Output as _items_ItemsList__Output } from '../items/ItemsList';

export interface ItemsServiceClient extends grpc.Client {
  Create(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Create(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Create(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Create(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  create(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  create(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  create(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  create(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  
  Delete(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Delete(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Delete(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Delete(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  delete(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  delete(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  delete(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  delete(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  
  Get(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Get(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Get(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Get(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  get(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  get(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  get(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  get(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  
  GetAll(argument: _items_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  GetAll(argument: _items_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  GetAll(argument: _items_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  GetAll(argument: _items_Empty, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  getAll(argument: _items_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  getAll(argument: _items_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  getAll(argument: _items_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  getAll(argument: _items_Empty, callback: grpc.requestCallback<_items_ItemsList__Output>): grpc.ClientUnaryCall;
  
  Update(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Update(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Update(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  Update(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  update(argument: _items_Item, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  update(argument: _items_Item, metadata: grpc.Metadata, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  update(argument: _items_Item, options: grpc.CallOptions, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  update(argument: _items_Item, callback: grpc.requestCallback<_items_Item__Output>): grpc.ClientUnaryCall;
  
}

export interface ItemsServiceHandlers extends grpc.UntypedServiceImplementation {
  Create: grpc.handleUnaryCall<_items_Item__Output, _items_Item>;
  
  Delete: grpc.handleUnaryCall<_items_Item__Output, _items_Item>;
  
  Get: grpc.handleUnaryCall<_items_Item__Output, _items_Item>;
  
  GetAll: grpc.handleUnaryCall<_items_Empty__Output, _items_ItemsList>;
  
  Update: grpc.handleUnaryCall<_items_Item__Output, _items_Item>;
  
}

export interface ItemsServiceDefinition extends grpc.ServiceDefinition {
  Create: MethodDefinition<_items_Item, _items_Item, _items_Item__Output, _items_Item__Output>
  Delete: MethodDefinition<_items_Item, _items_Item, _items_Item__Output, _items_Item__Output>
  Get: MethodDefinition<_items_Item, _items_Item, _items_Item__Output, _items_Item__Output>
  GetAll: MethodDefinition<_items_Empty, _items_ItemsList, _items_Empty__Output, _items_ItemsList__Output>
  Update: MethodDefinition<_items_Item, _items_Item, _items_Item__Output, _items_Item__Output>
}
