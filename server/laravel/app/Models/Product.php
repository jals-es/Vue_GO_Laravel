<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Product extends Model
{
    use HasFactory;
    protected $table = 'prods';

    public function types()
    {
        return $this->hasMany(ProductType::class, 'id_prod' , 'id');
    }
    public function extras()
    {
        return $this->hasMany(ProductExtras::class, 'id_prod' , 'id');
    }

}
