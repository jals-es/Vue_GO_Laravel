<?php

namespace App\Http\Controllers\Api\V1;

use JWTAuth;
use Tymon\JWTAuth\Exceptions\JWTException;
use App\Http\Controllers\Controller;
use App\Models\Bar;
use App\Models\User;
use Illuminate\Http\Request;
use App\Http\Controllers\Api\V2\AuthController;
use Exception;
use Illuminate\Support\Facades\DB;

class BarController extends Controller
{



    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function list(Request $request)
    {
        $bar = new Bar();
        return response()->json(['data' => $bar->listBar($request->id_bar)], 201);
    }
    public function count(Request $request) {

        $bar = new Bar();
        return response()->json(['data' => $bar->listBar($request->id_bar)->count()], 201);
    }
    public function stats(Request $request) {
        $bar = new Bar();

        // return response()->json(['message' => $bar->getOrders($request->id_bar)], 201);


        return response()->json(['data' => $bar->getStats($request->id_bar)], 201);

        // return new Bar->monthlyOrders()
    }
    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\BarRequest  $request
     * @return \Illuminate\Http\Response
     */
    public function create(Request $request)
    {
        //
        return response()->json(['message' => 'Post create succesfully create'], 201);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Bar  $bar
     * @return \Illuminate\Http\Response
     */
    public function info(Bar $bar)
    {
        return response()->json(['message' => 'Post create succesfully info'], 201);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \App\Models\Bar  $bar
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, Bar $bar)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Bar  $bar
     * @return \Illuminate\Http\Response
     */
    public function destroy(Bar $bar)
    {
        //
    }
}
