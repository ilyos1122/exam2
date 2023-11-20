CREATE OR REPLACE FUNCTION order_products_tr() RETURNS TRIGGER LANGUAGE PLPGSQL
AS 
$$
DECLARE quantity_count BIGINT;
totalprice NUMERIC;
percentAmount NUMERIC;
product_price NUMERIC;
BEGIN
    
    NEW.price = (select price from product where id = NEW.product_id);
    NEW.sum = NEW.quantity * NEW.price;
    percentAmount = (NEW.quantity * NEW.price)/100 * NEW.discount_amount;
    IF NEW.discount_type LIKE 'fixed' THEN 
        NEW.sum = NEW.quantity * NEW.price - NEW.discount_amount;
    END IF; 
    IF NEW.discount_type LIKE 'percent' THEN
        New.sum = NEW.sum - percentAmount;
    END IF;
    SELECT SUM(quantity) FROM order_products WHERE order_id = NEW.order_id into quantity_count;
    SELECT SUM(sum) FROM order_products WHERE order_id = NEW.order_id into totalprice;
    UPDATE orders SET total_count = quantity_count, total_price = totalprice, updated_at = CURRENT_TIMESTAMP
    WHERE id = NEW.order_id;

    RETURN NEW;
END;
$$;

CREATE TRIGGER order_products_trigger
BEFORE INSERT ON order_products
FOR EACH ROW EXECUTE PROCEDURE
order_products_tr();





CREATE OR REPLACE FUNCTION orders_tr() RETURNS TRIGGER LANGUAGE PLPGSQL
AS 
$$
DECLARE
delvery_p NUMERIC;
BEGIN
    SELECT delivery_price FROM branches WHERE id = NEW.branch_id INTO delvery_p;
    NEW.delivery_price = delvery_p;
    RETURN NEW;
END;
$$;

CREATE TRIGGER orders_trigger
BEFORE INSERT ON orders
FOR EACH ROW EXECUTE PROCEDURE
orders_tr();

