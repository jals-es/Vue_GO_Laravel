<?php

namespace App\Models;

use Illuminate\Support\Facades\DB;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Models\Stats;

class Bar extends Model
{

    use HasFactory;
    public function getStats($id_bar) {
        $stats = new Stats();
        return $stats->monthlyOrders($id_bar);
    }
    public function getOrders() {
        DB::select(DB::raw('WITH p AS (SELECT sum(order_extras.price) AS total,order_extras.id ,order_prods.id AS prodsId, order_prods.id_order AS idOrder
        FROM order_prods
        LEFT JOIN order_extras
        ON order_prods.id = order_extras.id_order_prod
        GROUP BY order_prods.id
    ), p2 AS (SELECT sum(order_prods.price) AS totalProds, orders.id, order_prods.id AS prodsId
        FROM order_prods
        LEFT JOIN orders
        ON order_prods.id_order = orders.id
        WHERE id_bar = 12345
        GROUP BY orders.id)
    SELECT p2.id, sum(coalesce(p2.totalProds+p.total, p2.totalProds, p.total)) AS price
    FROM p
    LEFT JOIN p2
    ON p.idOrder = p2.id
    WHERE p2.id IS NOT null
    GROUP BY p2.id'));
    }
}
