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
        try {
            if (app(AuthController::class)->check($request)) {
                $bars = DB::table('bars')
                ->get();
                return response()->json([
                    'success' => true,
                    'message' => 'Bar list',
                    'data' => $bars
                ], 200);
            } else {
                return response()->json([
                    'success' => false,
                    'message' => 'Forbiddenn'
                ], 403);
            }
        } catch (Exception $e) {
            return response()->json([
                'success' => false,
                'message' => 'Forbidden'
            ], 403);
        }
        return response()->json(['message' => 'Post create succesfully' . $token], 201);
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\BarRequest  $request
     * @return \Illuminate\Http\Response
     */
    public function create(BarRequest $request)
    {
        //
        return response()->json(['message' => 'Post create succesfully'], 201);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Bar  $bar
     * @return \Illuminate\Http\Response
     */
    public function info(Bar $bar)
    {
        return response()->json(['message' => 'Post create succesfully'], 201);
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
