<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\V1\BarController;
use App\Http\Controllers\Api\V1\IncidenceController;
use App\Http\Controllers\Api\V1\ChartsController;
use App\Http\Controllers\Api\V1\ProductsController;
use App\Http\Controllers\Api\V1\MessagesController;
use App\Http\Controllers\Api\V2\AuthController;
use Illuminate\Support\Facades\Mail;
use App\Mail\AdminNotification;


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
// Route::get('send_test_email', function(){
//     dispatch(new App\Jobs\SendEmailJob(env('ADMINMAIL')));
// });
    Route::name('api.')->group(function () {
        Route::middleware([SuperAdmin::class])->group(function(){
            Route::get('/check', function () {
                return response()->json(['status' =>'success', 'data' => 'authorized'], 201);
            });
            Route::name('message.')->group(function () {
                Route::post('message', [MessagesController::class, 'store'])->name('store');
                Route::get('message', [MessagesController::class, 'fetch'])->name('fetch');
            });
            Route::name('incidence.')->group(function () {
                Route::post('incidence', [IncidenceController::class, 'store'])->name('store');
                Route::get('incidence', [IncidenceController::class, 'list'])->name('list');
                Route::patch('incidence', [IncidenceController::class, 'close'])->name('close');
                Route::delete('incidence/{id}', [IncidenceController::class, 'delete'])->name('delete');
            });
            Route::name('bars.')->group(function () {
                Route::get('bars', [BarController::class, 'list'])->name('list');
                Route::get('bars/count', [BarController::class, 'count'])->name('count');
                Route::get('bars/stats', [BarController::class, 'stats'])->name('stats');
                Route::get('bars/stats/{slug}', [BarController::class, 'stats'])->name('stats');
                Route::get('bars/search/{slug}', [BarController::class, 'list'])->name('search');
                Route::get('bars/info/{slug}', [BarController::class, 'info'])->name('info');
                Route::post('bars', [BarController::class, 'create'])->name('create');
                Route::get('token', [BarController::class, 'list'])->name('list');
                Route::get('bars/orders/{id_bar}', [BarController::class, 'orders'])->name('orders');
                Route::get('bars/products/{id_bar}', [ProductsController::class, 'getProducts'])->name('getProducts');
            });
            Route::get('charts/first', [ChartsController::class, 'getFirstChartData'])->name('firstChart');
            Route::get('charts/second', [ChartsController::class, 'getSecondChartData'])->name('secondChart');
            Route::get('userData', [AuthController::class, 'userData'])->name('userData');
        });
        Route::post('auth', [AuthController::class, 'auth'])->name('auth');

    });
// });
