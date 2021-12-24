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
            $data[$key]->own = self::checkId($data[$key]->owner);
            $data[$key]->photos = $ind->photos;
        }
        return $data;
    }
    public function update()
    {
        $inc = Incidence::where('id', request()->id)->first();
        if ($inc->status == 0) {
            $inc->closeIncidence();
            $inc->update(['closer' => self::getId(), 'fix' => request()->incidenceFix]);
        }
    }
    // public function deleteIncidence()
    // {
    //     // if (Gate::allows('isOwner', Incidence::find() request()->)) {
    //     //     echo 'Allowed';
    //     //   } else {
    //     //     echo 'Not Allowed';
    //     //   }
    // }
    public function createIncidence()
    {
        try {
            $uuid = Str::uuid();
            $newIncidence = Incidence::create([
                'id' => $uuid,
                'name' => request()->name,
                'status' => 1,
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
