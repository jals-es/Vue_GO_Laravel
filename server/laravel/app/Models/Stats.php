<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class Stats extends Model
{
    use HasFactory;
    public function monthlyOrders($id_bar = '%') {
        if ($id_bar == null) {
            $id_bar = 'id_bar';
        }
        // return $id_bar;
        return DB::table('orders')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-2592000'))
            ->whereRaw('id_bar = ' . $id_bar)
            ->get()
            ->count();
    }
    public function dailyOrders($id_bar = '%') {
        if ($id_bar == null) {
            $id_bar = 'id_bar';
        }
        return DB::table('orders')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-86400'))
            ->whereRaw('id_bar = ' . $id_bar)
            ->get()
            ->count();
    }
    public function dailyActiveBars($id_bar = '%') {
        if ($id_bar == null) {
            $id_bar = 'id_bar';
        }
        return DB::table('orders')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-86400'))
            ->whereRaw('id_bar = ' . $id_bar)
            ->groupBy('id_bar')
            ->get()
            ->count();
    }
    public function getOrders($id_bar = 'id_bar') {
        if ($id_bar == null) {
            $id_bar = 'id_bar';
        }
        return DB::select(DB::raw('WITH p AS (SELECT sum(order_extras.price) AS total,order_extras.id ,order_prods.id AS prodsId, order_prods.id_order AS idOrder
        FROM order_prods
        LEFT JOIN order_extras
        ON order_prods.id = order_extras.id_order_prod
        GROUP BY order_prods.id
    ), p2 AS (SELECT sum(order_prods.price) AS totalProds, orders.id, order_prods.id AS prodsId
        FROM order_prods
        LEFT JOIN orders
        ON order_prods.id_order = orders.id
        WHERE id_bar = '. $id_bar .'
        GROUP BY orders.id)
    SELECT p2.id, sum(coalesce(p2.totalProds+p.total, p2.totalProds, p.total)) AS price
    FROM p
    LEFT JOIN p2
    ON p.idOrder = p2.id
    WHERE p2.id IS NOT null
    GROUP BY p2.id'));
    }
}
