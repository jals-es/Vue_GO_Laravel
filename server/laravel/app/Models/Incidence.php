<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Incidence extends Model
{
    use HasFactory;
    public $incrementing = false; // Si no laravel automaticament pensa que es una clau primaria AI i fica 0, 1, 2 en vegada del verdader id
    protected $table = 'incidences';
    protected $fillable = ['name','descr','id','owner','status'];

    public function photos()
    {
        return $this->hasMany(IncidencePhoto::class, 'id_incidence' , 'id');
    }
    public function owner()
    {
        return $this->belongsTo(User::class, 'id' , 'owner');
    }
    public function closeIncidence()
    {
        $this->status = 1;
        $this->save();
    }
    public function getName()
    {
        return $this->status == 1 ? 'CERRADO: '.$this->name : $this->name;
    }
    public static function boot() {
        parent::boot();
        //Delete photos on cascade :)
        static::deleting(function($incidence) { // Before delete method
             $incidence->photos()->delete();
        });
    }
}
