<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use App\Models\Message;
use App\Events\NewMessage;
use App\Traits\UserTrait;

class MessagesController extends Controller
{
    use UserTrait;
    public function fetch() {
        return response()->json(['status' =>'success', 'data' => Message::all()], 201);
    }
    public function store() {


        $newMessage = Message::create([
            'content' => request()->content,
            // 'author' => self::getId(),
            'author' => "antoni",

        ]);
        event(new NewMessage($newMessage));
        broadcast(new NewMessage($newMessage))->toOthers();

        return ['status' => 'Message Sent!'];
        // broadcast(new MessageSent($user, $message))->toOthers();
        // return response()->json(['status' =>'success', 'data' => $this->incidenceRepository->update()], 201);
    }
}
