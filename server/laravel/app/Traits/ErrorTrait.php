<?php

namespace App\Traits;
use Illuminate\Http\Response;


trait ErrorTrait {
	public static function throwError($message = "Internal Server Error")
	{
		return response()->json([
            'status' => false,
            'message' => $message,
            'data' => null,
        ], Response::HTTP_INTERNAL_SERVER_ERROR);
	}
}
