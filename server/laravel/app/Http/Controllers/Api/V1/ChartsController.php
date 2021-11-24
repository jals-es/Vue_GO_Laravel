<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Models\Charts;


class ChartsController extends Controller
{

    public function getChart1Data(Request $request) {
        $modelCharts = new Charts();
        return response()->json(['message' => $modelCharts->getChartData()], 201);

        // return new Bar->monthlyOrders()
    }
}
