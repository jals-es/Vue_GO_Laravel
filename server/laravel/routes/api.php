<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\V1\BarController;
use App\Http\Controllers\Api\V1\ChartsController;
use App\Http\Controllers\Api\V2\AuthController;



/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

// Route::apiResource('v1/bars', App\Http\Controllers\Api\V1\BarController::class)->middleware('api');

// Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
//     return $request->user();
// });
// Route::middleware(['cors'])->group(function () {
    Route::name('api.')->group(function () {
        Route::middleware([SuperAdmin::class])->group(function(){
            Route::name('bars.')->group(function () {

                Route::get('bars', [BarController::class, 'list'])->name('list');
                Route::get('bars/count', [BarController::class, 'count'])->name('count');
                Route::get('bars/stats', [BarController::class, 'stats'])->name('stats');

                Route::get('bars/{slug}', [BarController::class, 'info'])->name('info');
                Route::post('bars', [BarController::class, 'create'])->name('create');

                Route::get('token', [BarController::class, 'list'])->name('list');
            });
            Route::get('charts', [ChartsController::class, 'getChart1Data'])->name('chart1');
            Route::get('userData', [AuthController::class, 'userData'])->name('userData');

        });
        Route::post('auth', [AuthController::class, 'auth'])->name('auth');

    });
// });
