<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Incidence extends Model
{
    use HasFactory;

    protected $table = 'incidences';
    protected $fillable = ['name','descr','id','owner'];

    public function photos()
    {
        return $this->hasMany(IncidencePhoto::class, 'id_incidence' , 'id');
    }

}
