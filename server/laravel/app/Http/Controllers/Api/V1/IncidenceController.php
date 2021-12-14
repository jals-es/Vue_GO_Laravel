<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Support\Str;
use App\Http\Requests\IncidenceRequest;
use App\Repositories\IncidenceRepository;
class IncidenceController extends Controller
{
    public function __construct(IncidenceRepository $IncidenceRepository)
    {
        $this->incidenceRepository = $IncidenceRepository;
    }
    public function store(IncidenceRequest $request) {
        return response()->json(['status' =>'success', 'data' => $this->incidenceRepository->createIncidence($request)], 201);
        // return response()->json(['status' =>'success', 'data' => Str::uuid() ], 201);

    }

}
