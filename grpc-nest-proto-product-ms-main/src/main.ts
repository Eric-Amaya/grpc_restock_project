import { INestMicroservice, ValidationPipe, Logger, BadRequestException } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { Transport } from '@nestjs/microservices';
import { join } from 'path';
import { AppModule } from './app.module';
import { protobufPackage } from './product/product.pb';

async function bootstrap() {
  const logger = new Logger('Bootstrap');
  const app: INestMicroservice = await NestFactory.createMicroservice(AppModule, {
    transport: Transport.GRPC,
    options: {
      url: '0.0.0.0:5053',
      package: protobufPackage,
      protoPath: join(__dirname, '../node_modules/grpc-nest-proto/proto/product.proto'),
    },
  });

  app.useGlobalPipes(new ValidationPipe({
    whitelist: true,
    transform: true,
    exceptionFactory: (errors) => {
      logger.error('Validation errors:', errors);
      return new BadRequestException(errors);
    }
  }));

  await app.listen();
  logger.log('Product microservice is listening');
}

bootstrap();
