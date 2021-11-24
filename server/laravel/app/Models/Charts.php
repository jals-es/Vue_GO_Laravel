<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;
use \stdClass;

class Charts extends Model
{
    use HasFactory;

    public function getChartData() {
        $data = DB::table('orders')
            ->selectRaw('bars.name, count(bars.id) AS orders')
            ->leftJoin('bars', 'bars.id', '=', 'orders.id_bar')
            ->groupBy('bars.id')
            ->get();
            $a = new stdClass();
            $a->names = array();
            $a->orders = array();
            foreach ($data as $element) {
                array_push($a->orders,$element->orders);
                array_push($a->names,$element->name);
            }
            return $a;
    }
}
