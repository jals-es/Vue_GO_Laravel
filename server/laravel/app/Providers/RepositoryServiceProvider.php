<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use App\Interfaces\ProductRepositoryInterface;
use App\Repositories\ProductRepository;
use App\Traits\UserTrait;
use Illuminate\Support\Facades\Gate;
use App\Repositories\IncidenceRepository;

class RepositoryServiceProvider extends ServiceProvider
{
    use UserTrait;
    /**
     * Register services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->bind(ProductRepositoryInterface::class, ProductRepository::class);
        $this->app->bind(IncidenceRepositoryInterface::class, IncidenceRepository::class);
    }

    /**
     * Bootstrap services.
     *
     * @return void
     */
    public function boot()
    {
        // Gate::define('isOwner', function($incidence) {
        //     // return "hola";
        //     return true;
        //     // return false;
        //     // return $incidence->owner == self::getId();
        //  });
    }
}
