<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\V1\BarController;
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
Route::name('api.')->group(function () {
    Route::name('bars.')->group(function () {
        Route::get('bars', [BarController::class, 'list'])->name('list');
        Route::get('bars/{slug}', [BarController::class, 'info'])->name('info');
        Route::post('bars', [BarController::class, 'create'])->name('create');
        // Route::put('articles/{slug}', [BarController::class, 'update'])->name('update');
        // Route::delete('articles/{slug}', [BarController::class, 'delete'])->name('delete');
        Route::post('auth', [AuthController::class, 'auth'])->name('auth');
        Route::get('token', [BarController::class, 'list'])->name('list');
    });
});
