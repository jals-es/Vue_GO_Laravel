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
}
