<?php
namespace App\Services;

use App\Models\Incidence;
use App\Models\IncidencePhoto;
use Illuminate\Support\Str;

class IncidenceService
{

    public function __construct(Incidence $incidence, IncidencePhoto $incidencePhoto)
    {
        $this->incidence = $incidence;
        $this->incidencePhoto = $incidencePhoto;
    }

    public function syncPhotos($incidence, array $photos): void
    {
        foreach ($photos as $photo) {
                $fileName = time().'_'.$photo->getClientOriginalName();
                $photo->move(public_path().'/uploads/', $fileName);
                IncidencePhoto::create([
                    'id' => Str::uuid(),
                    'name' => $fileName,
                    'id_incidence' => $incidence,
                    'path' => 'http://127.0.0.1:8000/uploads/'.$fileName,
                ]);
        }
    }
}
