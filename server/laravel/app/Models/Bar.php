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
    public function getOrders($id_bar) {
        $stats = new Stats();
        return $stats->getOrders($id_bar);
    }
}
