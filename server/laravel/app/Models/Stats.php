<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class Stats extends Model
{
    use HasFactory;
    public function monthlyOrders($slug = 'slug') {
        if ($slug == null) {
            $slug = 'slug';
        } else {
            $slug = '"' . $slug . '"';
        }
        // return $id_bar;
        return DB::table('orders')
            ->select('orders.*', 'bars.slug')
            ->leftJoin('bars','bars.id','=','id_bar')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-2592000'))
            ->whereRaw('slug = ' . $slug)
            ->get()
            ->count();
    }
    public function dailyOrders($slug = 'slug') {
        if ($slug == null) {
            $slug = 'slug';
        } else {
            $slug = '"' . $slug . '"';
        }
        return DB::table('orders')
            ->select('orders.*', 'bars.slug')
            ->leftJoin('bars','bars.id','=','id_bar')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-86400'))
            ->whereRaw('slug = ' . $slug)
            ->get()
            ->count();
    }
    public function dailyActiveBars($slug = 'slug') {
        if ($slug == null) {
            $slug = 'slug';
        } else {
            $slug = '"' . $slug . '"';
        }
        return DB::table('orders')
            ->select('orders.*', 'bars.slug')
            ->leftJoin('bars','bars.id','=','id_bar')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-86400'))
            ->whereRaw('slug = ' . $slug)
            ->groupBy('id_bar')
            ->get()
            ->count();
    }
    public function getOrders($slug = 'slug') {
        if ($slug == null) {
            $slug = 'slug';
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
        WHERE slug = '. $slug .'
        GROUP BY orders.id)
    SELECT p2.id, sum(coalesce(p2.totalProds+p.total, p2.totalProds, p.total)) AS price
    FROM p
    LEFT JOIN p2
    ON p.idOrder = p2.id
    WHERE p2.id IS NOT null
    GROUP BY p2.id'));
    }
}
