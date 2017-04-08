module Views.AlertList.Views exposing (view)

import Alerts.Types exposing (Alert, AlertGroup, Block)
import Views.AlertList.Types exposing (AlertListMsg(FilterAlerts))
import Views.AlertList.Filter exposing (silenced, receiver, matchers)
import Views.AlertList.FilterBar
import Html exposing (..)
import Html.Attributes exposing (..)
import Utils.Date
import Utils.Types exposing (Filter)
import Utils.Views exposing (..)
import Types exposing (Msg(MsgForAlertList, Noop, CreateSilenceFromAlert, UpdateFilter, AddLabel))


view : List AlertGroup -> Filter -> Html Msg -> Html Msg
view alertGroups filter errorHtml =
    let
        filteredGroups =
            receiver filter.receiver alertGroups
                |> silenced filter.showSilenced

        filterText =
            Maybe.withDefault "" filter.text

        alertHtml =
            if List.isEmpty filteredGroups then
                div [ class "mt2" ] [ text "no alerts found found" ]
            else
                ul
                    [ classList
                        [ ( "list", True )
                        , ( "pa0", True )
                        ]
                    ]
                    (List.map alertGroupView filteredGroups)
    in
        div []
            [ Views.AlertList.FilterBar.view filterText (Types.UpdateFilter filter) (MsgForAlertList FilterAlerts)
            , errorHtml
            , alertHtml
            ]


alertGroupView : AlertGroup -> Html Msg
alertGroupView alertGroup =
    li [ class "pa3 pa4-ns bb b--black-10" ]
        [ div [ class "mb3" ] (List.map alertHeader <| List.sort alertGroup.labels)
        , div [] (List.map blockView alertGroup.blocks)
        ]


blockView : Block -> Html Msg
blockView block =
    div [] (List.map alertView block.alerts)


alertView : Alert -> Html Msg
alertView alert =
    let
        id =
            Maybe.withDefault "" alert.silenceId

        b =
            if alert.silenced then
                buttonLink "fa-deaf" ("#/silences/" ++ id) "blue" Noop
            else
                buttonLink "fa-exclamation-triangle" "#/silences/new?keep=1" "dark-red" (CreateSilenceFromAlert alert)

        labels =
            List.filter (Tuple.first >> (/=) "alertname") alert.labels
    in
        div [ class "f6 mb3" ]
            [ div [ class "mb1" ]
                [ b
                , buttonLink "fa-bar-chart" alert.generatorUrl "black" Noop
                , p [ class "dib mr2" ] [ text <| Utils.Date.timeFormat alert.startsAt ]
                ]
            , div [ class "mb2 w-80-l w-100-m" ] (List.map labelButton labels)
            ]


labelButton : ( String, String ) -> Html Msg
labelButton (( key, value ) as label) =
    onClickMsgButton
        (key ++ "=" ++ value)
        (AddLabel Noop label)


alertHeader : ( String, String ) -> Html msg
alertHeader ( key, value ) =
    if key == "alertname" then
        b [ class "db f4 mr2 dark-red dib" ] [ text value ]
    else
        listButton "ph1 pv1" ( key, value )
