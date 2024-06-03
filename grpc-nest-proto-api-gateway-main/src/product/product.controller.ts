import { Controller, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { ProductServiceClient, PRODUCT_SERVICE_NAME, CreateProductRequest, FindOneRequest, DecreaseStockRequest } from './product.pb';
import { Body, Get, Param, Post } from '@nestjs/common';
import { Observable } from 'rxjs';

@Controller('products')
export class ProductController {
  private productService: ProductServiceClient;

  @Inject(PRODUCT_SERVICE_NAME)
  private readonly client: ClientGrpc;

  onModuleInit() {
    this.productService = this.client.getService<ProductServiceClient>(PRODUCT_SERVICE_NAME);
  }

  @Post()
  createProduct(@Body() request: CreateProductRequest): Observable<any> {
    return this.productService.createProduct(request);
  }

  @Get(':id')
  findOne(@Param('id') id: number): Observable<any> {
    const request: FindOneRequest = { id };
    return this.productService.findOne(request);
  }

  @Post('decrease-stock')
  decreaseStock(@Body() request: DecreaseStockRequest): Observable<any> {
    return this.productService.decreaseStock(request);
  }
}
