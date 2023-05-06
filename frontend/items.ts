import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { ItemsServiceClient as _items_ItemsServiceClient, ItemsServiceDefinition as _items_ItemsServiceDefinition } from './items/ItemsService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  items: {
    Empty: MessageTypeDefinition
    Item: MessageTypeDefinition
    ItemsList: MessageTypeDefinition
    ItemsService: SubtypeConstructor<typeof grpc.Client, _items_ItemsServiceClient> & { service: _items_ItemsServiceDefinition }
  }
}

