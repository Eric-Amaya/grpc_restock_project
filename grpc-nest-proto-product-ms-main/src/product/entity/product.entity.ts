import { BaseEntity, Column, Entity, OneToMany, PrimaryGeneratedColumn } from 'typeorm';
import { StockDecreaseLog } from './stock-decrease-log.entity';

@Entity()
export class Product extends BaseEntity {
    @PrimaryGeneratedColumn()
    public id!: number;

    @Column({ type: 'varchar', length: 255 })
    public name: string;

    @Column({ type: 'varchar', length: 255 })
    public sku: string;

    @Column({ type: 'varchar', length: 255 })
    public category: string;

    @Column({ type: 'text' })
    public description: string;

    @Column({ type: 'int' })
    public stock: number;

    @Column({ type: 'int' })
    public price: number;
    
    // OneToMany relationships

    @OneToMany(() => StockDecreaseLog, (StockDecreaseLog) => StockDecreaseLog.product)
    public stockDecreaseLogs: StockDecreaseLog[];
}