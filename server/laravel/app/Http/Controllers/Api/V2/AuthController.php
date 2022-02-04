<?php

namespace App\Http\Controllers\Api\V2;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Illuminate\Support\Facades\DB;
use App\Http\Requests\IncidenceRequest;

class AuthController extends Controller
{

    public function auth(Request $request)
    {
        //Get the info
        $isSuperAdmin = DB::table('superadmins')
            ->leftJoin('users', 'users.id', '=', 'superadmins.id_user')
            ->where('users.email', '=', $request->user)
            ->where('users.passwd', '=', $request->passwd)
            ->get();
        if (!$isSuperAdmin->isEmpty()) {
            //User exists

            if ($request->ip() == env("ALLOWED_SERVER")) {
                //Its from a safe server
                //Get the key
                $key = env("JWT_SUPERADMIN");
                //Create payload with the superuser id and 1 hour
                $payload = array(
                    "token_superadmin" => $isSuperAdmin[0]->id,
                    "iat" => now()->timestamp,
                    "iot" => now()->timestamp + 3600,
                );
                //Encode JWT
                $jwt = JWT::encode($payload, $key, 'HS256');
                //Return the response with the jwt
                return response()->json([
                    'success' => true,
                    'message' => 'Token authed successfully',
                    'data' => $jwt,
                ], Response::HTTP_OK);
            } else {
                //User can login but not from a safe server
                return response()->json([
                    'success' => false,
                    'message' => 'Forbidden',
                    'ip' => $request->ip(),
                    'acceped' => env("ALLOWED_SERVER")
                ], 403);
            }
        } else {
            //Invalid user credentials
            return response()->json([
                'success' => false,
                'message' => 'Invalid Credentials',
            ], 403);
        }
    }
    public function userData(Request $request) {
        $decoded = JWT::decode($request->bearerToken(), new Key(env("JWT_SUPERADMIN"), 'HS256'));
        $data = DB::table('users')
        ->select('name','email','photo')
        ->where('id', '=', $decoded->token_superadmin)
        ->get();
        return response()->json([
            'success' => true,
            'message' => $data,
        ], 200);
    }
    public function userId($jwt) {
        $decoded = JWT::decode($jwt, new Key(env("JWT_SUPERADMIN"), 'HS256'));
        $data = DB::table('users')
        ->select('id')
        ->where('id', '=', $decoded->token_superadmin)
        ->get();
        return $data;
    }
    public function check($request)
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
                    return true;
                } else {
                    return false;
                }
            } else {
                return false;
            }
        //Prevent JWT decode fail and other unexpected things
        } catch (\Throwable $th) {
            return false;
        }
    }
}
