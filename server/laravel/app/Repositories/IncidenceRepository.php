<?php

namespace App\Repositories;

use App\Models\Incidence;
use App\Traits\ErrorTrait;
use App\Http\Requests\IncidenceRequest;
use Illuminate\Support\Str;

class IncidenceRepository
{
    use ErrorTrait;
    public function list()
    {
        // $own = Incidence::where('owner', '=', app('App\Http\Controllers\Api\V2\AuthController')->userId(Str::substr(request()->header('Authorization', ''), 7))[0]->id)->get();
        // foreach ($own as $key => $inc) {
        //     $own[$key]->photos = $inc->photos;
        // }
        // $rest = Incidence::where('owner', '!=', app('App\Http\Controllers\Api\V2\AuthController')->userId(Str::substr(request()->header('Authorization', ''), 7))[0]->id)->get();
        // foreach ($rest as $key => $inc) {
        //     $rest[$key]->photos = $inc->photos;
        // }
        $data = Incidence::all();
        foreach ($data as $key => $ind) {
            $data[$key]->own = $data[$key]->owner == app('App\Http\Controllers\Api\V2\AuthController')->userId(Str::substr(request()->header('Authorization', ''), 7))[0]->id ? true : false;
            $data[$key]->photos = $ind->photos;
        }
        return $data;
    }
    public function createIncidence()
    {
        try {
            $uuid = Str::uuid();
            $newIncidence = Incidence::create([
                'id' => $uuid,
                'name' => request()->name,
                'status' => 1,
                'owner' => app('App\Http\Controllers\Api\V2\AuthController')->userId(Str::substr(request()->header('Authorization', ''), 7))[0]->id,
                'closer' => null,
                'descr' => request()->descr,
                'fix' => null,
            ]);
            if ($newIncidence) {
                return $uuid;
            }
        } catch (\Exception $e) {
            return self::throwError($e);
        }
    }
}
