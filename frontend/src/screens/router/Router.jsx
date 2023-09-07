import React from "react";
import Home from "../home/Home.jsx"

import { BrowserRouter, Route, Routes } from "react-router-dom";
import PeerTable from "../tables/peer.table/PeerTable.jsx";
import TaskTable from "../tables/task.table/TaskTable.jsx";
import ChecksTable from "../tables/checks.table/ChecksTable.jsx";
import P2pTable from "../tables/p2p.table/P2pTable.jsx";
import VerterTable from "../tables/verter.table/VerterTable.jsx";
import TransferredPointsTable from "../tables/transferred_points.table/TransferredPointsTable.jsx";
import FriendsTable from "../tables/friends.table/FriendsTable.jsx";
import RecommendationsTable from "../tables/recommendations.table/RecommendationsTable.jsx";
import XpTable from "../tables/xp.table/XpTable.jsx";
import TimeTrackingTable from "../tables/time_tracking.table/TimeTrackingTable.jsx";
import ErrorNotFound from "../error/not_found.error.jsx/NotFound.jsx";
import TransferredPoints from "../functions/transferred_points.functions/TransferredPoints.jsx";
import XpTask from "../functions/xp_task.functions/XpTask.jsx";
import PeersDontLeave from "../functions/peers_dont_leave.functions/PeersDontLeave.jsx";
import SuccessFailureChecks from "../functions/success_failure_checks.functions/SuccessFailureChecks.jsx";
import PointsChangeV1 from "../functions/points_change_v1.functions/PointsChangeV1.jsx";
import PointsChangeV2 from "../functions/points_change_v2.functions/PointsChangeV2.jsx";
import OftenTaskPerDay from "../functions/often_task_per_day.functions/OftenTaskPerDay.jsx";
import LastP2pDuration from "../functions/last_p2p_duration.functions/LastP2pDuration.jsx";
import ListLastExPeer from "../functions/list_last_ex_peer.functions/ListLastExPeer.jsx";
import PeersForP2p from "../functions/peers_for_p2p.functions/PeersForP2p.jsx";
import StatisticBlock from "../functions/statistic_block.functions/StatisticBlock.jsx";
import MostFriendly from "../functions/most_friendly.functions/MostFriendly.jsx";
import SuccessAtBirthday from "../functions/success_at_birthday.functions/SuccessAtBirthday.jsx";
import PeerXpSum from "../functions/peer_xp_sum.functions/PeerXpSum.jsx";
import PassOneTwo from "../functions/pass_one_two.functions/PassOneTwo.jsx";
import PreviousTasks from "../functions/previous_tasks.functions/PreviousTasks.jsx";
import SuccessfulDays from "../functions/successful_days.functions/SuccessfulDays.jsx";
import PeerMostTasks from "../functions/peer_most_tasks.functions/PeerMostTasks.jsx";
import PeerMostXp from "../functions/peer_most_xp.functions/PeerMostXp.jsx";
import MaxTimeDate from "../functions/max_time_date.functions/MaxTimeDate.jsx";
import TimePeerByTime from "../functions/time_peer_by_time.functions/TimePeerByTime.jsx";
import EnterPeerByDay from "../functions/enter_peer_by_day.functions/EnterPeerByDay.jsx";
import LastFeastCame from "../functions/last_feast_came.functions/LastFeastCame.jsx";
import MoreThenTimePeer from "../functions/more_then_time_peer.functions/MoreThenTimePeer.jsx";
import EarlyEntries from "../functions/early_entries.functions/EarlyEntries.jsx";
import RawQuery from "../raw_query/RawQuery.jsx";

const Router = () => {
  return <BrowserRouter>
    <Routes>
      <Route path='/' element={<Home />} />
      <Route path='peers/' element={<PeerTable />} />
      <Route path='tasks/' element={<TaskTable />} />
      <Route path='checks/' element={<ChecksTable />} />
      <Route path='p2p/' element={<P2pTable />} />
      <Route path='verter/' element={<VerterTable />} />
      <Route path='transferred_points/' element={<TransferredPointsTable />} />
      <Route path='friends/' element={<FriendsTable />} />
      <Route path='recommendations/' element={<RecommendationsTable />} />
      <Route path='xp/' element={<XpTable />} />
      <Route path='time_tracking/' element={<TimeTrackingTable />} />

      <Route path='fnc_transferred_points/' element={<TransferredPoints />} />
      <Route path='fnc_xp_task/' element={<XpTask />} />
      <Route path='fnc_peers_dont_leave/' element={<PeersDontLeave />} />
      <Route path='fnc_success_failure_checks/' element={<SuccessFailureChecks />} />
      <Route path='fnc_points_change_v1/' element={<PointsChangeV1 />} />
      <Route path='fnc_points_change_v2/' element={<PointsChangeV2 />} />
      <Route path='fnc_often_task_per_day/' element={<OftenTaskPerDay />} />
      <Route path='fnc_last_p2p_duration/' element={<LastP2pDuration />} />
      <Route path='fnc_list_last_ex_peer/' element={<ListLastExPeer />} />
      <Route path='fnc_peers_for_p2p/' element={<PeersForP2p />} />
      <Route path='fnc_statistic_block/' element={<StatisticBlock />} />
      <Route path='fnc_most_friendly/' element={<MostFriendly />} />
      <Route path='fnc_success_at_birthday/' element={<SuccessAtBirthday />} />
      <Route path='fnc_peer_xp_sum/' element={<PeerXpSum />} />
      <Route path='fnc_pass_one_two/' element={<PassOneTwo />} />
      <Route path='fnc_previous_tasks/' element={<PreviousTasks />} />
      <Route path='fnc_successful_days/' element={<SuccessfulDays />} />
      <Route path='fnc_peer_most_tasks/' element={<PeerMostTasks />} />
      <Route path='fnc_peer_most_xp/' element={<PeerMostXp />} />
      <Route path='fnc_max_time_date/' element={<MaxTimeDate />} />
      <Route path='fnc_time_peer_by_time/' element={<TimePeerByTime />} />
      <Route path='fnc_enter_peer_by_day/' element={<EnterPeerByDay />} />
      <Route path='fnc_last_feast_came/' element={<LastFeastCame />} />
      <Route path='fnc_more_then_time_peer/' element={<MoreThenTimePeer />} />
      <Route path='fnc_early_entries/' element={<EarlyEntries />} />

      <Route path='raw_query/' element={<RawQuery />} />

      <Route path='*' element={<ErrorNotFound />} />
    </Routes>
  </BrowserRouter>
}

export default Router;

