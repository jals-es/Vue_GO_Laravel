<?php

namespace App\Models;

use Illuminate\Support\Facades\DB;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Models\Stats;

class Bar extends Model
{
    use HasFactory;
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
    public function getStats($id_bar) {
        $stats = new Stats();
        return ["monthlyOrders" => $stats->monthlyOrders($id_bar), "dailyActiveBars" => $stats->dailyActiveBars($id_bar), "dailyOrders" => $stats->dailyOrders($id_bar)];
    }
    public function getOrders($id_bar) {
        $stats = new Stats();
        return $stats->getOrders($id_bar);
    }
}
