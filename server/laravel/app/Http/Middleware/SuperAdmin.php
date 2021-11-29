<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Illuminate\Support\Facades\DB;

class SuperAdmin
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle(Request $request, Closure $next)
    {
        try {
            //Decode the token
            $decoded = JWT::decode($request->bearerToken(), new Key(env("JWT_SUPERADMIN"), 'HS256'));
            //Get if there is other user with the same id in the superadmin table
            $isSuperAdmin = DB::table('superadmins')
                ->where('id_user', '=', $decoded->token_superadmin)
                ->get();
            //Check if the before query is empty
            if (!$isSuperAdmin->isEmpty()) {
                //Check if the token is timed out
                if (now()->timestamp < $decoded->iot) {
                    return $next($request);
                } else {
                    return false;
                }
            } else {
                return false;
            }
        //Prevent JWT decode fail and other unexpected things
        } catch (\Throwable $th) {
            return response()->json('Authentication Error');
        }


    }
}
