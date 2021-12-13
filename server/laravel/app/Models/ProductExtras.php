<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class ProductExtras extends Model
{
    use HasFactory;
    protected $table = 'prods_extras';


    public function product() {
        return $this->belongsTo(ProductType::class);
    }


}
