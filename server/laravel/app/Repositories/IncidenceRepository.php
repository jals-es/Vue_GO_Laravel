<?php

namespace App\Repositories;

use App\Models\Incidence;
use App\Traits\ErrorTrait;
use App\Http\Requests\IncidenceRequest;
use Illuminate\Support\Str;

class IncidenceRepository
{
    use ErrorTrait;
    public function createIncidence()
    {
        $header = request()->header('Authorization', '');

        if (Str::startsWith($header, 'Bearer ')) {
            $jwt = Str::substr($header, 7);
        }


        try {
            $uuid = Str::uuid();
            $newIncidence = Incidence::create([
                'id' => $uuid,
                'name' => request()->name,
                'status' => 1,
                'owner' => app('App\Http\Controllers\Api\V2\AuthController')->userId($jwt)[0]->id,
                'closer' => null,
                'descr' => request()->descr,
                'fix' => null,
            ]);
            if ($newIncidence) {
                return $uuid;
            }
        } catch (\Exception $e) {
            return self::throwError();
        }
    }
}
