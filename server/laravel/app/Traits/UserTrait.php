<?php

namespace App\Traits;

use App\Models\Incidence;
use Illuminate\Http\Response;
use Illuminate\Support\Str;


trait UserTrait {
	public static function getId()
	{
		return app('App\Http\Controllers\Api\V2\AuthController')->userId(Str::substr(request()->header('Authorization', ''), 7))[0]->id;
	}
    public static function checkId($id)
    {
        return $id == self::getId() ? true : false;
    }

}
