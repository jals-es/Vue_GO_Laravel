<?php

namespace App\Http\Controllers\Api\V1;

use JWTAuth;
use Tymon\JWTAuth\Exceptions\JWTException;
use App\Http\Controllers\Controller;
use App\Models\Bar;
use App\Models\Orders;
use Illuminate\Http\Request;
use App\Http\Controllers\Api\V2\AuthController;
use Exception;
use Illuminate\Support\Facades\DB;

class BarController extends Controller
{
    public function list(Request $request)
    {
        $bar = new Bar();
        return response()->json(['data' => $bar->listBar("%",$request->slug)], 201);
    }
    public function count(Request $request) {

        $bar = new Bar();
        return response()->json(['data' => $bar->listBar($request->id_bar)->count()], 201);
    }
    public function stats(Request $request) {
        $bar = new Bar();
        return response()->json(['data' => $bar->getStats($request->slug)], 201);
    }
    public function info(Request $request) {
        $bar = new Bar();
        return response()->json(['data' => $bar->getInfo($request->slug)], 201);
    }
    public function orders(Request $request) {
        $bar = new Bar();
        return response()->json(['data' => $bar->getOrders($request->id_bar)], 201);
    }
}
