<?php

namespace App\Repositories;

use App\Models\Incidence;
use App\Traits\ErrorTrait;
use App\Traits\UserTrait;
use App\Http\Requests\IncidenceRequest;
use Illuminate\Support\Str;
use Illuminate\Support\Facades\Gate;

class IncidenceRepository
{
    use ErrorTrait;
    use UserTrait;
    public function list()
    {
        $data = Incidence::all();
        foreach ($data as $key => $ind) {
            $data[$key]->name = $data[$key]->getName();
            $data[$key]->own = self::checkId($data[$key]->owner);
            $data[$key]->photos = $ind->photos;
        }
        return $data;
    }
    public function update()
    {
        $inc = Incidence::where('id', request()->id);
        if ($inc->first()->status == 0) {
            $inc->first()->closeIncidence();
            $inc->update(['closer' => self::getId(), 'fix' => request()->incidenceFix]);
            return $inc;
        }
    }

    public function createIncidence()
    {
        try {
            $uuid = Str::uuid();
            $newIncidence = Incidence::create([
                'id' => $uuid,
                'name' => request()->name,
                'status' => 0,
                'owner' => self::getId(),
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
