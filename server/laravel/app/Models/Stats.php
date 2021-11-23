<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class Stats extends Model
{
    use HasFactory;
    public function monthlyOrders($id_bar = '%') {
        return DB::table('orders')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-2592000'))
            ->where('id_bar', '=', $id_bar)
            ->get()
            ->count();
    }
    public function dailyOrders($id_bar = '%') {
        return DB::table('orders')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-86400'))
            ->where('id_bar', '=', $id_bar)
            ->get()
            ->count();
    }
    public function dailyActiveBars($id_bar = '%') {
        return DB::table('orders')
            ->whereRaw(DB::raw('UNIX_TIMESTAMP(date) > UNIX_TIMESTAMP()-86400'))
            ->where('id_bar', '=', $id_bar)
            ->groupBy('id_bar')
            ->get()
            ->count();
    }
}
