<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class IncidencePhoto extends Model
{
    use HasFactory;
    protected $table = 'incidence_photos';

    public function incidence() {
        return $this->belongsTo(Incidence::class);
    }
}
