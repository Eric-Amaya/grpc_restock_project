import { Controller, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { OrderServiceClient, ORDER_SERVICE_NAME, CreateOrderRequest, GetOrderRequest, GetAllOrdersRequest } from './order.pb';
import { Observable } from 'rxjs';
import { Body, Get, Param, Post } from '@nestjs/common';

@Controller('orders')
export class OrderController {
  private orderService: OrderServiceClient;

  @Inject(ORDER_SERVICE_NAME)
  private readonly client: ClientGrpc;

  onModuleInit() {
    this.orderService = this.client.getService<OrderServiceClient>(ORDER_SERVICE_NAME);
  }

  @Post()
  createOrder(@Body() request: CreateOrderRequest): Observable<any> {
    return this.orderService.createOrder(request);
  }

  @Get(':id')
  getOrder(@Param('id') id: number): Observable<any> {
    const request: GetOrderRequest = { orderId: id };
    return this.orderService.getOrder(request);
  }

  @Get()
  getAllOrders(): Observable<any> {
    const request: GetAllOrdersRequest = {};
    return this.orderService.getAllOrders(request);
  }
}
