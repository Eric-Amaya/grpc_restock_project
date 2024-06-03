import { IsNotEmpty, IsNumber, IsString, IsOptional } from 'class-validator';
import { CreateProductRequest, DecreaseStockRequest, FindOneRequest } from './product.pb';

export class FindOneRequestDto implements FindOneRequest {
  @IsNumber({ allowInfinity: false, allowNaN: false })
  public readonly id: number;
}

export class CreateProductRequestDto implements CreateProductRequest {
  @IsString()
  @IsNotEmpty()
  public readonly name: string;

  @IsString()
  @IsNotEmpty()
  public readonly sku: string;

  @IsString()
  @IsNotEmpty()
  public readonly category: string = 'default category';  // Default value

  @IsString()
  @IsNotEmpty()
  public readonly description: string = 'default description';  // Default value

  @IsNumber({ allowInfinity: false, allowNaN: false })
  public readonly price: number;

  @IsNumber({ allowInfinity: false, allowNaN: false })
  public readonly stock: number;
}

export class DecreaseStockRequestDto implements DecreaseStockRequest {
  @IsNumber({ allowInfinity: false, allowNaN: false })
  public readonly id: number;

  @IsOptional()
  @IsNumber({ allowInfinity: false })
  public readonly orderId: number;
}
