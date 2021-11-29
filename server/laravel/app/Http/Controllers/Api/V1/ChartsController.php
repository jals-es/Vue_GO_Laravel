<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Models\Charts;


class ChartsController extends Controller
{

    public function getFirstChartData() {
        $modelCharts = new Charts();
        return response()->json(['status' =>'success', 'data' => $modelCharts->getFirstChartData()], 201);

        // return new Bar->monthlyOrders()
    }
}
