<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use App\Services\IncidenceService;
use App\Http\Requests\IncidenceRequest;
use Illuminate\Http\Request;

use App\Repositories\IncidenceRepository;
class IncidenceController extends Controller
{

    public function __construct(IncidenceRepository $IncidenceRepository, IncidenceService $incidenceService)
    {
        $this->incidenceService = $incidenceService;
        $this->incidenceRepository = $IncidenceRepository;
    }
    public function list() {
        return response()->json(['status' =>'success', 'data' => $this->incidenceRepository->list()], 201);
    }
    public function store(Request $request) {
        $request->validate([
            'file' => 'required',
            'file.*' => 'mimes:jpeg,jpg,png|max:2048'
          ]);
        $createIncidence = $this->incidenceRepository->createIncidence();
        $this->incidenceService->syncPhotos($createIncidence, $request->file('file'));
        return response()->json(['status' =>'success', 'data' => 'succes created'], 201);

    }

}
