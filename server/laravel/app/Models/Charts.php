<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;
use \stdClass;

class Charts extends Model
{
    use HasFactory;

    public function getFirstChartData() {
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
    public function getSecondChartData() {

        $data = DB::table('orders')
                ->selectRaw('DAYNAME(date) as day, count(*) as total')
                ->groupBy('day')
                ->get();
                $a = new stdClass();
                $a->days = array();
                $a->total = array();
                foreach ($data as $element) {
                    array_push($a->days,$element->day);
                    array_push($a->total,$element->total);
                }
                return $a;

    }

    // SELECT DAYNAME(date) as day, count(*) as total FROM orders
// GROUP BY day
}
