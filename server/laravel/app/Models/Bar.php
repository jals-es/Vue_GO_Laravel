<?php

namespace App\Models;

use Illuminate\Support\Facades\DB;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Models\Stats;
use App\Models\Orders;


class Bar extends Model
{
    use HasFactory;
    public function getInfo($slug) {
        return DB::table('bars')
        ->selectRaw("bars.*,users.name as ownerName")
        ->leftJoin('users', 'users.id', '=', 'bars.owner')
        ->where('bars.slug', '=', $slug)
        // ->where('bars.name', 'LIKE', $name)
        ->get();
    }
    public function listBar($id, $name = "%") {
        if ($id == null) {
            $id = '%';
        }
        return DB::table('bars')
        ->selectRaw("bars.*,users.name as ownerName")
        ->leftJoin('users', 'users.id', '=', 'bars.owner')
        ->where('bars.id', 'LIKE', $id)
        ->where('bars.name', 'LIKE', $name)
        ->get();
    }
    public function getStats($slug = 'slug') {
        $stats = new Stats();
        return ["monthlyOrders" => $stats->monthlyOrders($slug), "dailyActiveBars" => $stats->dailyActiveBars($slug), "dailyOrders" => $stats->dailyOrders($slug)];
    }
    public function getOrders($id_bar) {
        $orders = new Orders();
        return $orders->getOrders($id_bar);
    }
}
