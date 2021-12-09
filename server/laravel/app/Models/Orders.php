<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class Orders extends Model
{
    use HasFactory;
    public function getOrders($slug = 'slug') {
        if ($slug == null) {
            $slug = '12345';
        }
        return DB::select(DB::raw('WITH p AS (SELECT sum(order_extras.price) AS total,order_extras.id ,order_prods.id AS prodsId, order_prods.id_order AS idOrder
        FROM order_prods
        LEFT JOIN order_extras
        ON order_prods.id = order_extras.id_order_prod
        GROUP BY order_prods.id
    ), p2 AS (SELECT sum(order_prods.price) AS totalProds, orders.*, order_prods.id AS prodsId
        FROM order_prods
        LEFT JOIN orders
        ON order_prods.id_order = orders.id
        WHERE orders.id_bar LIKE ' . $slug . '
        GROUP BY orders.id)
        SELECT p2.id_table,p2.date,u.name,b.slug, sum(coalesce(p2.totalProds+p.total, p2.totalProds, p.total)) AS price
        FROM p
        LEFT JOIN p2
        ON p.idOrder = p2.id
        LEFT JOIN users u
        ON u.id = p2.id_user
        LEFT JOIN bars b
        ON b.id = p2.id_bar
        WHERE p2.id IS NOT null
        GROUP BY p2.id'));
    }
}
